// Package home implements the Tsuki home/landing page.
package home

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/msgs"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// menuItem represents a single navigation option on the home page.
type menuItem struct {
	label       string
	description string
	page        msgs.PageID
}

// Model is the home page model.
type Model struct {
	cursor int
	items  []menuItem
	width  int
	height int
}

// New returns an initialised home page model.
func New() Model {
	return Model{
		items: []menuItem{
			{
				label:       "📝  Notes",
				description: "Write and browse your lunar notes",
				page:        msgs.Notes,
			},
			{
				label:       "⚙   Settings",
				description: "Configure your preferences",
				page:        msgs.Settings,
			},
		},
	}
}

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd { return nil }

// Update handles key events for the home page.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter", " ":
			return m, navigateTo(m.items[m.cursor].page)
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the home page.
func (m Model) View() string {
	var b strings.Builder

	// Banner
	b.WriteString(styles.Title.Render(styles.MoonBanner))
	b.WriteString("\n\n")

	// Tagline
	b.WriteString(styles.Subtitle.Render("  Your personal lunar journal"))
	b.WriteString("\n\n")

	// Divider
	b.WriteString(styles.Divider.Render("  " + strings.Repeat("·", 34)))
	b.WriteString("\n\n")

	// Menu items
	for i, item := range m.items {
		cursor := "   "
		labelStyle := styles.Normal
		descStyle := styles.Muted

		if i == m.cursor {
			cursor = " ▶ "
			labelStyle = styles.Selected
			descStyle = lipgloss.NewStyle().Foreground(styles.ColorSecondary)
		}

		b.WriteString(cursor + labelStyle.Render(item.label) + "\n")
		b.WriteString("     " + descStyle.Render(item.description) + "\n\n")
	}

	// Help bar
	b.WriteString("\n")
	b.WriteString(styles.Help.Render("  ↑/↓  navigate   enter  select   q  quit"))

	content := b.String()
	if m.width > 0 && m.height > 0 {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
	}
	return content
}

// SetSize stores the terminal dimensions for layout purposes.
func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
}

func navigateTo(page msgs.PageID) tea.Cmd {
	return func() tea.Msg { return msgs.NavigateMsg{To: page} }
}
