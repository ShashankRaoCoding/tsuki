package app

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
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
	runErr  error
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
	switch msg := msg.(type) {
	case cliLaunchResult:
		l.runErr = msg.err
		return nil
	}

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
	parts := []string{styles.Subtitle.Render("Select a CLI and press Enter to launch it.")}
	if l.loadErr != nil {
		parts = append(parts, styles.ErrorStyle.Render(fmt.Sprintf("config load error: %v", l.loadErr)))
	}
	if l.runErr != nil {
		parts = append(parts, styles.ErrorStyle.Render(fmt.Sprintf("launch error: %v", l.runErr)))
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

type cliLaunchResult struct {
	err error
}

func runCLI(cfg cliConfig) tea.Cmd {
	return func() tea.Msg {
		parts := strings.Fields(cfg.Syntax)
		if len(parts) == 0 {
			return cliLaunchResult{err: fmt.Errorf("empty syntax for %s", cfg.Name)}
		}
		cmd := exec.Command(parts[0], parts[1:]...) // #nosec G204 -- command is loaded from local config
		execCmd := tea.ExecProcess(cmd, func(err error) tea.Msg {
			return cliLaunchResult{err: err}
		})
		return execCmd()
	}
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
