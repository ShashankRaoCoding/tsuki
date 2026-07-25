// Package app implements the root Bubble Tea model that orchestrates page routing.
package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/msgs"
	"github.com/ShashankRaoCoding/tsuki/internal/pages/home"
	"github.com/ShashankRaoCoding/tsuki/internal/pages/notes"
	"github.com/ShashankRaoCoding/tsuki/internal/pages/settings"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// App is the root model.  It owns all page models and delegates to the active one.
type App struct {
	page     msgs.PageID
	home     home.Model
	notes    notes.Model
	settings settings.Model
	width    int
	height   int
}

// New returns an initialised App.
func New() App {
	return App{
		page:     msgs.Home,
		home:     home.New(),
		notes:    notes.New(),
		settings: settings.New(),
	}
}

// Init satisfies tea.Model.
func (a App) Init() tea.Cmd { return nil }

// Update handles global messages and delegates the rest to the active page.
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.home.SetSize(msg.Width, msg.Height)
		a.notes.SetSize(msg.Width, msg.Height)
		a.settings.SetSize(msg.Width, msg.Height)
		return a, nil

	case msgs.NavigateMsg:
		a.page = msg.To
		return a, nil
	}

	var cmd tea.Cmd
	switch a.page {
	case msgs.Home:
		a.home, cmd = a.home.Update(msg)
	case msgs.Notes:
		a.notes, cmd = a.notes.Update(msg)
	case msgs.Settings:
		a.settings, cmd = a.settings.Update(msg)
	}
	return a, cmd
}

// View renders the active page inside a consistent chrome (top bar + content area).
func (a App) View() string {
	if a.width == 0 {
		// Terminal size not yet known — render page without chrome.
		return a.pageView()
	}

	// Top status bar.
	bar := a.statusBar()

	// Page content with padding.
	content := a.pageView()
	if a.page != msgs.Home {
		// Non-home pages are padded from the top-left rather than centred.
		content = lipgloss.NewStyle().
			PaddingTop(1).
			PaddingLeft(4).
			Render(content)
	}

	return bar + "\n" + content
}

// statusBar returns a full-width top bar showing the current page.
func (a App) statusBar() string {
	label := map[msgs.PageID]string{
		msgs.Home:     "Home",
		msgs.Notes:    "Notes",
		msgs.Settings: "Settings",
	}[a.page]

	left := styles.Title.Render("🌙 Tsuki")
	right := styles.Muted.Render(label)

	gap := a.width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}

	bar := lipgloss.NewStyle().
		Background(lipgloss.Color("#1E1E2E")).
		Foreground(styles.ColorText).
		Width(a.width).
		Render(left + strings.Repeat(" ", gap) + right)

	return bar
}

// pageView returns the View() output of the currently active page.
func (a App) pageView() string {
	switch a.page {
	case msgs.Home:
		return a.home.View()
	case msgs.Notes:
		return a.notes.View()
	case msgs.Settings:
		return a.settings.View()
	}
	return ""
}
