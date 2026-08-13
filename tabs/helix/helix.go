package helix

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/shell"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

const maxPreviewLines = 120

type helixExitMsg struct {
	err error
}

// EditorWidget manages selected file preview and launches Helix.
type EditorWidget struct {
	width        int
	height       int
	openedPath   string
	previewLines []string
	status       string
}

// NewEditorWidget returns an editor widget.
func NewEditorWidget() shell.Widget {
	return &EditorWidget{
		status: "Select a file from the sidebar and press enter to launch Helix.",
	}
}

// Init satisfies shell.Widget.
func (w *EditorWidget) Init() tea.Cmd { return nil }

// Update satisfies shell.Widget.
func (w *EditorWidget) Update(msg tea.Msg) (shell.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
	case shell.OpenFileMsg:
		w.openedPath = msg.Path
		w.loadPreview(msg.Path)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return w, w.launchHelix()
		}
	case helixExitMsg:
		if msg.err != nil {
			w.status = fmt.Sprintf("Helix exited: %v", msg.err)
		} else {
			w.status = "Helix closed."
		}
	}
	return w, nil
}

// View satisfies shell.Widget.
func (w *EditorWidget) View() string {
	lines := []string{
		styles.Title.Render("Helix"),
		"",
		styles.Muted.Render("Choose a file in the left pane, then press enter here to launch Helix."),
	}
	if w.openedPath != "" {
		lines = append(lines, "", "file: "+w.openedPath)
	}
	if w.status != "" {
		lines = append(lines, "", styles.Muted.Render(w.status))
	}
	if len(w.previewLines) > 0 {
		lines = append(lines, "", styles.Subtitle.Render("Preview"))
		lines = append(lines, w.previewLines...)
	}
	lines = append(lines, "", styles.Help.Render("enter launch helix • alt+left files • ctrl+q quit"))
	return lipgloss.NewStyle().Padding(1).Render(strings.Join(lines, "\n"))
}

func (w *EditorWidget) launchHelix() tea.Cmd {
	if strings.TrimSpace(w.openedPath) == "" {
		w.status = "No file selected."
		return nil
	}
	w.status = "Launching Helix..."
	path := w.openedPath
	cmd := exec.Command("hx", path)
	return tea.ExecProcess(cmd, func(err error) tea.Msg { return helixExitMsg{err: err} })
}

func (w *EditorWidget) loadPreview(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		w.previewLines = nil
		w.status = fmt.Sprintf("Cannot read file: %v", err)
		return
	}

	if bytes.IndexByte(data, 0) >= 0 {
		w.previewLines = []string{styles.Muted.Render("(binary file preview unavailable)")}
		w.status = "Ready to launch Helix."
		return
	}

	textLines := strings.Split(string(data), "\n")
	if len(textLines) > maxPreviewLines {
		w.previewLines = append([]string{}, textLines[:maxPreviewLines]...)
		w.previewLines = append(w.previewLines, styles.Muted.Render(fmt.Sprintf("… %d more lines", len(textLines)-maxPreviewLines)))
	} else {
		w.previewLines = textLines
	}

	w.status = "Ready to launch Helix."
}

// FilesWidget lists repository files for Helix.
type FilesWidget struct {
	width   int
	height  int
	cursor  int
	root    string
	entries []string
	status  string
}

// NewFilesWidget returns the left-hand file widget.
func NewFilesWidget() shell.Widget {
	root, err := os.Getwd()
	w := &FilesWidget{root: root}
	if err != nil {
		w.status = fmt.Sprintf("Cannot resolve working directory: %v", err)
		return w
	}
	w.refresh()
	return w
}

func (w *FilesWidget) refresh() {
	entries, err := collectFiles(w.root)
	if err != nil {
		w.status = fmt.Sprintf("Cannot read files: %v", err)
		w.entries = nil
		w.cursor = 0
		return
	}
	w.entries = entries
	if len(w.entries) == 0 {
		w.status = "No files found."
		w.cursor = 0
		return
	}
	if w.cursor >= len(w.entries) {
		w.cursor = len(w.entries) - 1
	}
	w.status = fmt.Sprintf("%d file(s)", len(w.entries))
}

func collectFiles(root string) ([]string, error) {
	entries := make([]string, 0, 128)
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			name := d.Name()
			if name == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if !d.Type().IsRegular() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		entries = append(entries, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(entries)
	return entries, nil
}

// Init satisfies shell.Widget.
func (w *FilesWidget) Init() tea.Cmd { return nil }

// Update satisfies shell.Widget.
func (w *FilesWidget) Update(msg tea.Msg) (shell.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			w.refresh()
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}
		case "down", "j":
			if w.cursor < len(w.entries)-1 {
				w.cursor++
			}
		case "enter":
			if len(w.entries) == 0 {
				return w, nil
			}
			path := w.entries[w.cursor]
			return w, func() tea.Msg { return shell.OpenFileMsg{Path: filepath.Join(w.root, path)} }
		}
	}
	return w, nil
}

// View satisfies shell.Widget.
func (w *FilesWidget) View() string {
	lines := []string{
		styles.Title.Render("Files"),
		"",
		styles.Muted.Render("Repository files"),
		"",
	}
	if w.status != "" {
		lines = append(lines, styles.Muted.Render(w.status), "")
	}
	if len(w.entries) == 0 {
		lines = append(lines, styles.Muted.Render("No files available."))
	}
	for i, entry := range w.entries {
		prefix := "  "
		style := styles.Normal
		if i == w.cursor {
			prefix = "→ "
			style = styles.Selected
		}
		lines = append(lines, prefix+style.Render(entry))
	}
	lines = append(lines, "", styles.Help.Render("↑/↓ move • enter open • r refresh • alt+right helix"))
	return lipgloss.NewStyle().Padding(1).Render(strings.Join(lines, "\n"))
}
