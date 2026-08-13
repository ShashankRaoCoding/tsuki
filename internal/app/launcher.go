package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

type cliConfig struct {
	Name        string
	Syntax      string
	Description string
}

type launchCLIRequest struct {
	Config cliConfig
}

type launcherModel struct {
	list    list.Model
	configs []cliConfig
	loadErr error
}

func newLauncherModel(configs []cliConfig, loadErr error) *launcherModel {
	items := make([]list.Item, 0, len(configs))
	for _, cfg := range configs {
		items = append(items, cliItem{cfg: cfg})
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Launcher"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(true)
	l.Styles.Title = styles.Title

	return &launcherModel{
		list:    l,
		configs: configs,
		loadErr: loadErr,
	}
}

func (l *launcherModel) Title() string {
	return "Launcher"
}

func (l *launcherModel) Init() tea.Cmd {
	return nil
}

func (l *launcherModel) SetSize(width, height int) {
	if height < 1 {
		height = 1
	}
	l.list.SetSize(width, height)
}

func (l *launcherModel) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	l.list, cmd = l.list.Update(msg)

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok || keyMsg.String() != "enter" {
		return cmd
	}

	selected, ok := l.list.SelectedItem().(cliItem)
	if !ok {
		return cmd
	}

	return tea.Batch(cmd, func() tea.Msg {
		return launchCLIRequest{Config: selected.cfg}
	})
}

func (l *launcherModel) View() string {
	parts := []string{styles.Subtitle.Render("Select a CLI and press Enter to open a tab.")}
	if l.loadErr != nil {
		parts = append(parts, styles.ErrorStyle.Render(fmt.Sprintf("config load error: %v", l.loadErr)))
	}
	if len(l.configs) == 0 {
		parts = append(parts, styles.Muted.Render("No CLI definitions found in ./CLIs"))
	}
	parts = append(parts, l.list.View())
	return strings.Join(parts, "\n\n")
}

type cliItem struct {
	cfg cliConfig
}

func (c cliItem) Title() string       { return c.cfg.Name }
func (c cliItem) Description() string { return c.cfg.Description }
func (c cliItem) FilterValue() string { return c.cfg.Name + " " + c.cfg.Syntax }

type execResult struct {
	prompt string
	output string
	err    error
}

type cliTab struct {
	cfg     cliConfig
	input   textinput.Model
	history []string
	height  int
	running bool
}

func newCLITab(cfg cliConfig) *cliTab {
	input := textinput.New()
	input.Placeholder = fmt.Sprintf("%s <args>", strings.TrimSpace(cfg.Syntax))
	input.Focus()
	input.CharLimit = 512

	promptLabel := strings.ToLower(strings.TrimSpace(cfg.Name))
	if promptLabel == "" {
		promptLabel = "cli"
	}
	input.Prompt = promptLabel + "> "

	return &cliTab{
		cfg:   cfg,
		input: input,
	}
}

func (c *cliTab) Title() string {
	return c.cfg.Name
}

func (c *cliTab) Init() tea.Cmd {
	return textinput.Blink
}

func (c *cliTab) SetSize(_ int, height int) {
	c.height = height
}

func runCommand(syntax, command string) tea.Cmd {
	syntax = strings.TrimSpace(syntax)
	command = strings.TrimSpace(command)
	if syntax == "" {
		return func() tea.Msg {
			return execResult{prompt: command, err: fmt.Errorf("no command configured")}
		}
	}

	prompt := fmt.Sprintf("%s %s", syntax, command)
	return func() tea.Msg {
		args := strings.Fields(command)
		cmd := exec.Command(syntax, args...) // #nosec G204 -- user-provided shell arguments
		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		err := cmd.Run()
		return execResult{prompt: prompt, output: buf.String(), err: err}
	}
}

func (c *cliTab) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case execResult:
		c.running = false
		c.history = append(c.history, styles.Muted.Render("$ "+msg.prompt))
		if msg.output != "" {
			c.history = append(c.history, strings.TrimRight(msg.output, "\n"))
		}
		if msg.err != nil {
			c.history = append(c.history, styles.ErrorStyle.Render(msg.err.Error()))
		}
		return nil

	case tea.KeyMsg:
		if msg.String() == "enter" && !c.running {
			command := strings.TrimSpace(c.input.Value())
			c.input.SetValue("")
			if command == "" {
				return nil
			}
			c.running = true
			return runCommand(c.cfg.Syntax, command)
		}
	}

	var cmd tea.Cmd
	c.input, cmd = c.input.Update(msg)
	return cmd
}

func (c *cliTab) View() string {
	header := styles.Title.Render(c.cfg.Name) + "  " + styles.Muted.Render("("+c.cfg.Description+")")
	body := styles.Muted.Render("No commands yet.")
	if len(c.history) > 0 {
		start := 0
		visible := c.height - 4
		if visible < 1 {
			visible = 1
		}
		if len(c.history) > visible {
			start = len(c.history) - visible
		}
		body = strings.Join(c.history[start:], "\n")
	}

	var prompt string
	if c.running {
		prompt = styles.Muted.Render("running…")
	} else {
		prompt = c.input.View()
	}
	return strings.Join([]string{header, "", body, "", prompt}, "\n")
}

func loadCLIConfigs(dir string) ([]cliConfig, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	configs := make([]cliConfig, 0)
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		content, readErr := os.ReadFile(filepath.Join(dir, entry.Name()))
		if readErr != nil {
			return nil, readErr
		}

		var raw struct {
			Name        string `json:"name"`
			Syntax      string `json:"syntax"`
			Description string `json:"description"`
			DescAlt     string `json:"Description"`
		}
		if unmarshalErr := json.Unmarshal(content, &raw); unmarshalErr != nil {
			return nil, fmt.Errorf("%s: %w", entry.Name(), unmarshalErr)
		}

		description := strings.TrimSpace(raw.Description)
		if description == "" {
			description = strings.TrimSpace(raw.DescAlt)
		}

		configs = append(configs, cliConfig{
			Name:        strings.TrimSpace(raw.Name),
			Syntax:      strings.TrimSpace(raw.Syntax),
			Description: description,
		})
	}

	sort.Slice(configs, func(i, j int) bool {
		return strings.ToLower(configs[i].Name) < strings.ToLower(configs[j].Name)
	})

	return configs, nil
}
