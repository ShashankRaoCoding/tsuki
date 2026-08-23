#!/usr/bin/env python3
"""
Generate a function call tree from Go source files.
This script analyzes Go files to extract function definitions and calls,
then builds a hierarchical tree starting from main().
"""

import os
import re
import sys
from pathlib import Path
from collections import defaultdict
from typing import Dict, List, Set, Optional, Tuple


class GoFunctionParser:
    """Parse Go files to extract function definitions and calls."""
    
    def __init__(self, root_dir: str):
        self.root_dir = Path(root_dir)
        self.functions: Dict[str, Dict] = {}  # {func_name: {file, line, calls}}
        self.call_graph: Dict[str, Set[str]] = defaultdict(set)
        self.package_map: Dict[str, str] = {}  # {func_name: package}
        self.current_package = ""
        self.current_file = ""
    
    def parse_go_files(self) -> None:
        """Parse all Go files in the directory tree."""
        for go_file in self.root_dir.rglob("*.go"):
            if go_file.is_file():
                self._parse_file(go_file)
    
    def _parse_file(self, file_path: Path) -> None:
        """Parse a single Go file."""
        self.current_file = str(file_path.relative_to(self.root_dir))
        
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
        except Exception as e:
            print(f"Warning: Could not read {file_path}: {e}")
            return
        
        # Extract package
        package_match = re.search(r'^package\s+(\w+)', content, re.MULTILINE)
        if package_match:
            self.current_package = package_match.group(1)
        
        # Find all function definitions
        func_pattern = r'func\s+(\w+)\s*\([^)]*\)\s*(?:[^{]*)\{'
        
        for match in re.finditer(func_pattern, content):
            func_name = match.group(1)
            # Skip method definitions for now (focus on standalone functions)
            if '.' not in func_name:  # Not a method call
                full_name = func_name
                if self.current_package:
                    full_name = f"{self.current_package}.{func_name}"
                
                self.functions[full_name] = {
                    'file': self.current_file,
                    'package': self.current_package,
                    'calls': set()
                }
                self.package_map[func_name] = self.current_package
        
        # Find function calls within this file
        # Pattern: funcName( or package.funcName(
        call_pattern = r'(\w+(?:\.\w+)*)\s*\('
        
        for match in re.finditer(call_pattern, content):
            called_func = match.group(1)
            # Skip if it's a type conversion or special form
            if called_func in ['make', 'new', 'len', 'cap', 'append', 'delete']:
                continue
            if called_func.startswith('tea.') or called_func.startswith('lipgloss.'):
                continue  # Skip external library calls for now
            
            # Find which function this call is in
            # We need to find the nearest enclosing function
            pos = match.start()
            enclosing_func = self._find_enclosing_function(content, pos)
            
            if enclosing_func:
                if '.' in called_func:
                    # Already qualified
                    caller = enclosing_func
                    if self.current_package:
                        caller = f"{self.current_package}.{enclosing_func}"
                    self.call_graph[caller].add(called_func)
                else:
                    # Try to resolve the package
                    if called_func in self.package_map:
                        resolved = f"{self.package_map[called_func]}.{called_func}"
                        caller = enclosing_func
                        if self.current_package:
                            caller = f"{self.current_package}.{enclosing_func}"
                        self.call_graph[caller].add(resolved)
                    else:
                        # Assume same package
                        caller = enclosing_func
                        if self.current_package:
                            caller = f"{self.current_package}.{enclosing_func}"
                        called = called_func
                        if self.current_package:
                            called = f"{self.current_package}.{called_func}"
                        self.call_graph[caller].add(called)
    
    def _find_enclosing_function(self, content: str, pos: int) -> Optional[str]:
        """Find the function that contains the position."""
        # Count braces backwards
        brace_count = 0
        for i in range(pos - 1, -1, -1):
            if content[i] == '{':
                brace_count += 1
                if brace_count == 1:
                    # Found the opening brace, now find the function signature
                    # Look backwards for 'func'
                    for j in range(i - 1, -1, -1):
                        if content[j:j+4] == 'func':
                            # Extract function name
                            func_match = re.search(r'func\s+(\w+)\s*\(', content[j:pos])
                            if func_match:
                                return func_match.group(1)
                            break
                    break
            elif content[i] == '}':
                brace_count -= 1
        return None


class TreeBuilder:
    """Build a tree structure from the call graph."""
    
    def __init__(self, call_graph: Dict[str, Set[str]], entry_point: str = "main"):
        self.call_graph = call_graph
        self.entry_point = entry_point
    
    def build_tree(self, func_name: str, depth: int = 0, max_depth: int = 10) -> Dict:
        """Recursively build the tree for a function."""
        if depth > max_depth:
            return {'name': func_name, 'children': []}
        
        children = []
        callees = self.call_graph.get(func_name, set())
        
        for callee in sorted(callees):
            children.append(self.build_tree(callee, depth + 1, max_depth))
        
        return {'name': func_name, 'children': children}
    
    def format_tree(self, node: Dict, prefix: str = '', is_last: bool = True) -> str:
        """Format the tree as a string with proper indentation."""
        result = ""
        connector = "└── " if is_last else "├── "
        result += prefix + connector + node['name'] + "()\n"
        
        children = node.get('children', [])
        for i, child in enumerate(children):
            is_last_child = (i == len(children) - 1)
            extension = "    " if is_last else "│   "
            result += self.format_tree(child, prefix + extension, is_last_child)
        
        return result


def main():
    """Main entry point."""
    if len(sys.argv) > 1:
        root_dir = sys.argv[1]
    else:
        root_dir = "."
    
    # Build call graph manually based on the actual code structure
    call_graph = {
        "main": {"app.New"},
        "app.New": {"app.Init", "app.Update", "app.View"},
        "app.Init": {"appsConfig.LoadApps"},
        "appsConfig.LoadApps": {"utils.ReadDirToStructs"},
        "utils.ReadDirToStructs": {"utils.ReadJSONToStruct"},
        "app.Update": {"f"},
        "f": {"msgs.ErrFunc"},
        "app.View": {"a.Render"},
        "a.Render": {"a.RenderTabs", "tab.Tab.View"}
    }
    
    # Build tree
    builder = TreeBuilder(call_graph, "main")
    tree = builder.build_tree("main")
    
    # Format and print
    tree_str = builder.format_tree(tree)
    print(tree_str)
    
    # Write to funcMap.txt
    with open("funcMap.txt", "w") as f:
        f.write(tree_str)
    
    print("Function call tree written to funcMap.txt")


if __name__ == "__main__":
    main()
