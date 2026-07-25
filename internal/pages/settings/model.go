// Package settings implements the Tsuki settings page.
package settings

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ShashankRaoCoding/tsuki/internal/msgs"
	"github.com/ShashankRaoCoding/tsuki/internal/page"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// setting represents a single toggleable preference.
type setting struct {
	key         string
	description string
	enabled     bool
}

// Model is the settings page model.
type Model struct {
	settings []setting
	cursor   int
	width    int
	height   int
}

// New returns an initialised settings page model.
func New() *Model {
	return &Model{
		settings: []setting{
			{
				key:         "Show timestamps",
				description: "Display the creation time beneath each note",
				enabled:     true,
			},
			{
				key:         "Compact view",
				description: "Reduce vertical spacing in the notes list",
				enabled:     false,
			},
			{
				key:         "Confirm deletes",
				description: "Ask for confirmation before removing a note",
				enabled:     false,
			},
		},
	}
}

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd { return nil }

// Update handles key events for the settings page.
func (m Model) Update(msg tea.Msg) (page.Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.settings)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.settings[m.cursor].enabled = !m.settings[m.cursor].enabled
		case "esc":
			return &m, navigateTo(msgs.Home)
		}
	}
	return &m, nil
}

// View renders the settings page.
func (m Model) View() string {
	var b strings.Builder

	// Header
	b.WriteString(styles.Title.Render("⚙   Settings") + "\n")
	b.WriteString(styles.Divider.Render(strings.Repeat("─", 44)) + "\n\n")

	// Settings list
	for i, s := range m.settings {
		prefix := "   "
		keyStyle := styles.Normal
		descStyle := styles.Muted

		if i == m.cursor {
			prefix = " ▶ "
			keyStyle = styles.Selected
			descStyle = styles.Subtitle
		}

		toggle := styles.ErrorStyle.Render("[ ]")
		if s.enabled {
			toggle = styles.SuccessStyle.Render("[✓]")
		}

		b.WriteString(prefix + toggle + "  " + keyStyle.Render(s.key) + "\n")
		b.WriteString("        " + descStyle.Render(s.description) + "\n\n")
	}

	// Help bar
	b.WriteString("\n")
	b.WriteString(styles.Help.Render("  ↑/↓  navigate   space/enter  toggle   esc  back"))

	return b.String()
}

// Title returns the display name of this page.
func (m Model) Title() string { return "Settings" }

// SetSize stores the terminal dimensions.
func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
}

func navigateTo(page msgs.PageID) tea.Cmd {
	return func() tea.Msg { return msgs.NavigateMsg{To: page} }
}
