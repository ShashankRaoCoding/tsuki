// Package app implements the root Bubble Tea model that orchestrates page routing.
package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/msgs"
	"github.com/ShashankRaoCoding/tsuki/internal/page"
	"github.com/ShashankRaoCoding/tsuki/internal/pages/home"
	"github.com/ShashankRaoCoding/tsuki/internal/pages/notes"
	"github.com/ShashankRaoCoding/tsuki/internal/pages/settings"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// App is the root model.  It owns all page models and delegates to the active one.
// To add a new page: implement page.Page, assign it a PageID in msgs, and append
// it to the pages slice in New() at the index matching that PageID.
type App struct {
	activePage msgs.PageID
	pages      []page.Page
	width      int
	height     int
}

// New returns an initialised App.
func New() App {
	return App{
		activePage: msgs.Home,
		pages: []page.Page{
			home.New(),     // index msgs.Home
			notes.New(),    // index msgs.Notes
			settings.New(), // index msgs.Settings
		},
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
		for _, p := range a.pages {
			p.SetSize(msg.Width, msg.Height)
		}
		return a, nil

	case msgs.NavigateMsg:
		a.activePage = msg.To
		return a, nil
	}

	var cmd tea.Cmd
	a.pages[a.activePage], cmd = a.pages[a.activePage].Update(msg)
	return a, cmd
}

// View renders the active page inside a consistent chrome (top bar + content area).
func (a App) View() string {
	if a.width == 0 {
		// Terminal size not yet known — render page without chrome.
		return a.pages[a.activePage].View()
	}

	// Top status bar.
	bar := a.statusBar()

	// Page content with padding.
	content := a.pages[a.activePage].View()
	if a.activePage != msgs.Home {
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
	left := styles.Title.Render("🌙 Tsuki")
	right := styles.Muted.Render(a.pages[a.activePage].Title())

	gap := a.width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}

	return lipgloss.NewStyle().
		Background(lipgloss.Color("#1E1E2E")).
		Foreground(styles.ColorText).
		Width(a.width).
		Render(left + strings.Repeat(" ", gap) + right)
}
