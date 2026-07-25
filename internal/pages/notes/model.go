// Package notes implements the Tsuki notes page with text input.
package notes

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ShashankRaoCoding/tsuki/internal/msgs"
	"github.com/ShashankRaoCoding/tsuki/internal/page"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// Note holds the content and creation time of a single note.
type Note struct {
	Content   string
	CreatedAt time.Time
}

// Model is the notes page model.
type Model struct {
	notes    []Note
	input    textinput.Model
	adding   bool
	cursor   int
	feedback string
	width    int
	height   int
}

// New returns an initialised notes page model.
func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Write your note here…"
	ti.CharLimit = 200
	ti.Width = 50
	ti.PromptStyle = styles.Title
	ti.TextStyle = styles.Normal

	return &Model{
		notes: []Note{
			{Content: "Welcome to Tsuki! 🌙", CreatedAt: time.Now()},
		},
		input: ti,
	}
}

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd { return nil }

// Update handles key events for the notes page.
func (m Model) Update(msg tea.Msg) (page.Page, tea.Cmd) {
	var cmd tea.Cmd

	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if m.adding {
			switch keyMsg.String() {
			case "esc":
				m.adding = false
				m.input.Reset()
				m.input.Blur()
				return &m, nil
			case "enter":
				val := strings.TrimSpace(m.input.Value())
				if val != "" {
					m.notes = append(m.notes, Note{
						Content:   val,
						CreatedAt: time.Now(),
					})
					m.cursor = len(m.notes) - 1
					m.feedback = "✓ Note saved"
				}
				m.input.Reset()
				m.input.Blur()
				m.adding = false
				return &m, nil
			}
			// Forward all other keys to the text input.
			m.input, cmd = m.input.Update(msg)
			return &m, cmd
		}

		// Navigation mode.
		switch keyMsg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			m.feedback = ""
		case "down", "j":
			if m.cursor < len(m.notes)-1 {
				m.cursor++
			}
			m.feedback = ""
		case "a", "n":
			m.adding = true
			m.feedback = ""
			cmd = m.input.Focus()
		case "d":
			if len(m.notes) > 0 {
				m.notes = append(m.notes[:m.cursor], m.notes[m.cursor+1:]...)
				if m.cursor >= len(m.notes) && m.cursor > 0 {
					m.cursor--
				}
				m.feedback = "Note deleted"
			}
		case "esc":
			return &m, navigateTo(msgs.Home)
		}
	}

	return &m, cmd
}

// View renders the notes page.
func (m Model) View() string {
	var b strings.Builder

	// Header
	header := styles.Title.Render("📝  Notes")
	count := styles.Muted.Render(fmt.Sprintf("  %d note(s)", len(m.notes)))
	b.WriteString(header + count + "\n")
	b.WriteString(styles.Divider.Render(strings.Repeat("─", 44)) + "\n\n")

	// Note list
	if len(m.notes) == 0 {
		b.WriteString(styles.Muted.Render("  No notes yet — press 'a' to add one.") + "\n")
	} else {
		for i, note := range m.notes {
			prefix := "   "
			noteStyle := styles.Normal
			timeStyle := styles.Muted

			if i == m.cursor {
				prefix = " ▶ "
				noteStyle = styles.Selected
				timeStyle = styles.Subtitle
			}

			b.WriteString(prefix + noteStyle.Render(note.Content) + "\n")
			b.WriteString("     " + timeStyle.Render(note.CreatedAt.Format("Jan 2  15:04")) + "\n\n")
		}
	}

	// Feedback line
	if m.feedback != "" {
		b.WriteString(styles.SuccessStyle.Render("  " + m.feedback) + "\n")
	}

	// Input area or help bar
	b.WriteString("\n")
	if m.adding {
		b.WriteString(styles.Title.Render("  New note") + "\n")
		b.WriteString("  " + m.input.View() + "\n\n")
		b.WriteString(styles.Help.Render("  enter  save   esc  cancel"))
	} else {
		b.WriteString(styles.Help.Render("  ↑/↓  navigate   a  add   d  delete   esc  back"))
	}

	return b.String()
}

// Title returns the display name of this page.
func (m Model) Title() string { return "Notes" }

// SetSize stores the terminal dimensions and adjusts the input width.
func (m *Model) SetSize(w, h int) {
	m.width = w
	m.height = h
	inputWidth := w - 12
	if inputWidth > 60 {
		inputWidth = 60
	}
	if inputWidth > 0 {
		m.input.Width = inputWidth
	}
}

func navigateTo(page msgs.PageID) tea.Cmd {
	return func() tea.Msg { return msgs.NavigateMsg{To: page} }
}
