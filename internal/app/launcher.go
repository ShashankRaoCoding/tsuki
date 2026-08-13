package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/creack/pty"

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

func (l *launcherModel) Close() {}

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

type cliOutputMsg struct {
	chunk string
}

type cliExitMsg struct {
	err error
}

type cliTab struct {
	cfg     cliConfig
	width   int
	height  int
	running bool
	lastErr error
	output  string

	cmd    *exec.Cmd
	cancel context.CancelFunc
	pty    *os.File
	events chan tea.Msg
}

func newCLITab(cfg cliConfig) *cliTab {
	return &cliTab{cfg: cfg}
}

func (c *cliTab) Title() string {
	return c.cfg.Name
}

func (c *cliTab) Init() tea.Cmd {
	return c.launch()
}

func (c *cliTab) SetSize(width, height int) {
	c.width = width
	c.height = height
	c.resizePTY()
}

func (c *cliTab) Close() {
	if c.cancel != nil {
		c.cancel()
	}
	if c.pty != nil {
		_ = c.pty.Close()
	}
}

func (c *cliTab) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case cliOutputMsg:
		c.appendOutput(msg.chunk)
		return c.waitForEvent()
	case cliExitMsg:
		c.running = false
		c.lastErr = msg.err
		c.releaseProcess()
		return nil
	case tea.KeyMsg:
		if msg.String() == "enter" && !c.running {
			return c.launch()
		}
		if c.running {
			c.forwardKey(msg)
		}
	}
	return nil
}

func (c *cliTab) View() string {
	header := styles.Title.Render(c.cfg.Name) + "  " + styles.Muted.Render("("+c.cfg.Description+")")
	status := styles.Muted.Render("Press Enter to relaunch.")
	if c.running {
		status = styles.SuccessStyle.Render("Running in panel")
	} else if c.lastErr != nil {
		status = styles.ErrorStyle.Render(fmt.Sprintf("Exited with error: %v", c.lastErr))
	} else {
		status = styles.SuccessStyle.Render("Exited successfully. Press Enter to relaunch.")
	}
	body := strings.TrimRight(c.panelBody(), "\n")
	if body == "" {
		body = styles.Muted.Render("No output yet.")
	}

	return strings.Join([]string{header, status, "", body}, "\n")
}

func (c *cliTab) launch() tea.Cmd {
	c.releaseProcess()

	parts := strings.Fields(c.cfg.Syntax)
	if len(parts) == 0 {
		c.running = false
		c.lastErr = fmt.Errorf("empty syntax for %s", c.cfg.Name)
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, parts[0], parts[1:]...) // #nosec G204 -- command is loaded from local config
	ptyFile, err := pty.Start(cmd)
	if err != nil {
		cancel()
		c.running = false
		c.lastErr = err
		return nil
	}

	c.running = true
	c.lastErr = nil
	c.output = ""
	c.cmd = cmd
	c.cancel = cancel
	c.pty = ptyFile
	c.events = make(chan tea.Msg, 64)
	c.resizePTY()

	go c.readOutput()
	go c.waitForExit()

	return c.waitForEvent()
}

func (c *cliTab) readOutput() {
	buf := make([]byte, 4096)
	for {
		n, err := c.pty.Read(buf)
		if n > 0 {
			c.events <- cliOutputMsg{chunk: string(buf[:n])}
		}
		if err != nil {
			if err != io.EOF {
				c.events <- cliOutputMsg{chunk: "\n[pty read error] " + err.Error() + "\n"}
			}
			return
		}
	}
}

func (c *cliTab) waitForExit() {
	err := c.cmd.Wait()
	c.events <- cliExitMsg{err: err}
	close(c.events)
}

func (c *cliTab) waitForEvent() tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-c.events
		if !ok {
			return nil
		}
		return msg
	}
}

func (c *cliTab) appendOutput(chunk string) {
	c.output += chunk
	const maxOutputSize = 200_000
	if len(c.output) > maxOutputSize {
		c.output = c.output[len(c.output)-maxOutputSize:]
	}
}

func (c *cliTab) panelBody() string {
	if c.output == "" {
		return ""
	}

	rows := c.height - 5
	if rows < 1 {
		rows = 1
	}

	output := strings.ReplaceAll(c.output, "\r\n", "\n")
	lines := strings.Split(output, "\n")
	if len(lines) > rows {
		lines = lines[len(lines)-rows:]
	}
	return strings.Join(lines, "\n")
}

func (c *cliTab) forwardKey(msg tea.KeyMsg) {
	if c.pty == nil {
		return
	}

	var payload string
	switch msg.String() {
	case "enter":
		payload = "\n"
	case "backspace":
		payload = "\x7f"
	case "esc":
		payload = "\x1b"
	case "up":
		payload = "\x1b[A"
	case "down":
		payload = "\x1b[B"
	default:
		if strings.HasPrefix(msg.String(), "ctrl+") && len(msg.String()) == len("ctrl+a") {
			ctrl := msg.String()[5]
			if ctrl >= 'a' && ctrl <= 'z' {
				payload = string(rune(ctrl - 'a' + 1))
			}
		}
		if payload == "" {
			payload = string(msg.Runes)
		}
	}

	if payload != "" {
		_, _ = c.pty.Write([]byte(payload))
	}
}

func (c *cliTab) resizePTY() {
	if c.pty == nil {
		return
	}

	rows := c.height - 5
	if rows < 1 {
		rows = 1
	}
	cols := c.width
	if cols < 1 {
		cols = 1
	}

	_ = pty.Setsize(c.pty, &pty.Winsize{
		Rows: uint16(rows),
		Cols: uint16(cols),
	})
}

func (c *cliTab) releaseProcess() {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
	if c.pty != nil {
		_ = c.pty.Close()
		c.pty = nil
	}
	c.cmd = nil
	c.events = nil
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
