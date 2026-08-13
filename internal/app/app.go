package app

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

type tab interface {
	Title() string
	Init() tea.Cmd
	Update(tea.Msg) tea.Cmd
	View() string
	SetSize(width, height int)
	Close()
}

type model struct {
	tabs   []tab
	focus  int
	width  int
	height int
}

// New creates the root Bubble Tea model.
func New() tea.Model {
	configs, loadErr := loadCLIConfigs("CLIs")
	launcher := newLauncherModel(configs, loadErr)

	m := &model{
		tabs: []tab{launcher},
	}

	return m
}

func (m *model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(m.tabs))
	for _, t := range m.tabs {
		cmds = append(cmds, t.Init())
	}
	return tea.Batch(cmds...)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		for _, t := range m.tabs {
			t.SetSize(m.width, m.bodyHeight())
		}
		return m, nil
	case launchCLIRequest:
		newTab := newCLITab(msg.Config)
		newTab.SetSize(m.width, m.bodyHeight())
		m.tabs = append(m.tabs, newTab)
		m.focus = len(m.tabs) - 1
		return m, newTab.Init()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			m.closeAllTabs()
			return m, tea.Quit
		case "ctrl+w":
			if m.focus > 0 && len(m.tabs) > 1 {
				m.tabs[m.focus].Close()
				m.tabs = append(m.tabs[:m.focus], m.tabs[m.focus+1:]...)
				if m.focus >= len(m.tabs) {
					m.focus = len(m.tabs) - 1
				}
			}
			return m, nil
		case "ctrl+t":
			m.focus = 0
			return m, nil
		case "tab", "right":
			_, isLauncher := m.tabs[m.focus].(*launcherModel)
			if isLauncher && len(m.tabs) > 0 {
				m.focus = (m.focus + 1) % len(m.tabs)
				return m, nil
			}
		case "shift+tab", "left":
			_, isLauncher := m.tabs[m.focus].(*launcherModel)
			if isLauncher && len(m.tabs) > 0 {
				m.focus = (m.focus - 1 + len(m.tabs)) % len(m.tabs)
				return m, nil
			}
		default:
			if idx, ok := parseTabIndex(msg.String()); ok && idx < len(m.tabs) {
				m.focus = idx
				return m, nil
			}
		}
	}

	if len(m.tabs) == 0 {
		return m, nil
	}

	cmd := m.tabs[m.focus].Update(msg)
	return m, cmd
}

func (m *model) View() string {
	if len(m.tabs) == 0 {
		return styles.ErrorStyle.Render("no tabs available")
	}

	tabHeaders := make([]string, 0, len(m.tabs))
	for i, t := range m.tabs {
		label := fmt.Sprintf("%d:%s", i+1, t.Title())
		if i == m.focus {
			tabHeaders = append(tabHeaders, styles.Selected.Render(label))
		} else {
			tabHeaders = append(tabHeaders, styles.Normal.Render(label))
		}
	}

	header := strings.Join(tabHeaders, styles.Divider.Render(" | "))
	help := styles.Help.Render("tab/shift+tab or ←/→: switch • ctrl+t: launcher • ctrl+w: close tab • ctrl+q: quit")

	return strings.Join([]string{header, "", m.tabs[m.focus].View(), "", help}, "\n")
}

func (m *model) bodyHeight() int {
	if m.height < 6 {
		return 1
	}
	return m.height - 5
}

func parseTabIndex(input string) (int, bool) {
	n, err := strconv.Atoi(input)
	if err != nil || n <= 0 {
		return 0, false
	}
	return n - 1, true
}

func (m *model) closeAllTabs() {
	for _, t := range m.tabs {
		t.Close()
	}
}
