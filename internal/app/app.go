package app

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

var (
	tabBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#11111B")).
			Foreground(styles.ColorText)

	activeTabStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Background(styles.ColorPrimary).
			Foreground(lipgloss.Color("#11111B")).
			Bold(true)

	inactiveTabStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Background(lipgloss.Color("#313244")).
				Foreground(styles.ColorText)
)

// App is the root Bubble Tea model for the tab shell.
type App struct {
	tabs       []Tab
	activeTab  int
	focusZone  FocusZone
	width      int
	height     int
	tabBarHeight int
	nextTabID  int
}

// New returns an initialised App with a single tab.
func New() App {
	return App{
		tabs:      []Tab{newDefaultTab(1)},
		activeTab: 0,
		focusZone: FocusTabRight,
		nextTabID: 2,
	}
}

// Init satisfies tea.Model.
func (a App) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0, len(a.tabs))
	for _, tab := range a.tabs {
		cmds = append(cmds, tab.Init())
	}
	return tea.Batch(cmds...)
}

// Update handles app-level keybindings and delegates the rest to the active tab.
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.tabBarHeight = lipgloss.Height(a.renderTabBar())
		if a.tabBarHeight < 1 {
			a.tabBarHeight = 1
		}
		return a, a.resizeTabs()
	case tea.KeyMsg:
		if next, cmd, handled := a.handleGlobalKey(msg); handled {
			return next, cmd
		}
	}

	if len(a.tabs) == 0 {
		return a, nil
	}

	var cmd tea.Cmd
	a.tabs[a.activeTab], a.focusZone, cmd = a.tabs[a.activeTab].Update(msg, a.focusZone)
	return a, cmd
}

// View renders the full-width tab bar and the active tab body.
func (a App) View() string {
	if len(a.tabs) == 0 {
		return ""
	}

	bar := a.renderTabBar()
	body := a.tabs[a.activeTab].View()
	if body == "" {
		return bar
	}
	return bar + "\n" + body
}

func (a App) handleGlobalKey(msg tea.KeyMsg) (App, tea.Cmd, bool) {
	if len(a.tabs) == 0 {
		return a, nil, false
	}

	switch msg.String() {
	case "ctrl+t":
		tab := newDefaultTab(a.nextTabID)
		a.nextTabID++
		a.tabs = append(a.tabs, tab)
		a.activeTab = len(a.tabs) - 1
		a.focusZone = defaultFocusZone(tab)

		cmds := []tea.Cmd{tab.Init()}
		if a.width > 0 || a.height > 0 {
			a.tabBarHeight = lipgloss.Height(a.renderTabBar())
			if a.tabBarHeight < 1 {
				a.tabBarHeight = 1
			}
			cmds = append(cmds, a.resizeTabs())
		}
		return a, tea.Batch(cmds...), true
	case "ctrl+w":
		if len(a.tabs) == 1 {
			a.tabs[0] = newDefaultTab(a.nextTabID)
			a.nextTabID++
			a.activeTab = 0
			a.focusZone = defaultFocusZone(a.tabs[0])

			cmds := []tea.Cmd{a.tabs[0].Init()}
			if a.width > 0 || a.height > 0 {
				cmds = append(cmds, a.resizeTabs())
			}
			return a, tea.Batch(cmds...), true
		}

		a.tabs = append(a.tabs[:a.activeTab], a.tabs[a.activeTab+1:]...)
		if a.activeTab >= len(a.tabs) {
			a.activeTab = len(a.tabs) - 1
		}
		a.focusZone = normalizeFocusZone(a.focusZone, a.tabs[a.activeTab])
		if a.width > 0 || a.height > 0 {
			a.tabBarHeight = lipgloss.Height(a.renderTabBar())
			if a.tabBarHeight < 1 {
				a.tabBarHeight = 1
			}
		}
		return a, nil, true
	case "ctrl+tab":
		a.activeTab = (a.activeTab + 1) % len(a.tabs)
		a.focusZone = normalizeFocusZone(a.focusZone, a.tabs[a.activeTab])
		return a, nil, true
	case "ctrl+shift+tab":
		a.activeTab--
		if a.activeTab < 0 {
			a.activeTab = len(a.tabs) - 1
		}
		a.focusZone = normalizeFocusZone(a.focusZone, a.tabs[a.activeTab])
		return a, nil, true
	default:
		return a, nil, false
	}
}

func (a *App) resizeTabs() tea.Cmd {
	bodyHeight := a.height - a.tabBarHeight
	if bodyHeight < 0 {
		bodyHeight = 0
	}

	cmds := make([]tea.Cmd, 0, len(a.tabs))
	for i := range a.tabs {
		tab, _, cmd := a.tabs[i].Update(tea.WindowSizeMsg{
			Width:  a.width,
			Height: bodyHeight,
		}, a.focusZone)
		a.tabs[i] = tab
		cmds = append(cmds, cmd)
	}

	if len(a.tabs) > 0 {
		a.focusZone = normalizeFocusZone(a.focusZone, a.tabs[a.activeTab])
	}

	return tea.Batch(cmds...)
}

func (a App) renderTabBar() string {
	if len(a.tabs) == 0 {
		return ""
	}

	tabs := make([]string, 0, len(a.tabs))
	for i, tab := range a.tabs {
		label := fmt.Sprintf("%d:%s", i+1, tab.title)
		if i == a.activeTab {
			tabs = append(tabs, activeTabStyle.Render(label))
			continue
		}
		tabs = append(tabs, inactiveTabStyle.Render(label))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Bottom, tabs...)
	if a.width <= 0 {
		return row
	}

	gap := a.width - lipgloss.Width(row)
	if gap < 0 {
		gap = 0
	}

	return tabBarStyle.Width(a.width).Render(row + strings.Repeat(" ", gap))
}

func defaultFocusZone(tab Tab) FocusZone {
	if tab.left != nil {
		return FocusTabRight
	}
	return FocusTabRight
}

func normalizeFocusZone(zone FocusZone, tab Tab) FocusZone {
	switch zone {
	case FocusGlobal:
		return FocusGlobal
	case FocusTabLeft:
		if tab.left != nil {
			return FocusTabLeft
		}
		return FocusTabRight
	case FocusTabRight:
		return FocusTabRight
	default:
		if tab.left != nil {
			return FocusTabLeft
		}
		return FocusTabRight
	}
}
