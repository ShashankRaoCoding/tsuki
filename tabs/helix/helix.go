package helix

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/shell"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// EditorWidget is a stub for a future PTY-backed Helix editor pane.
type EditorWidget struct {
	width      int
	height     int
	openedPath string
	lastKey    string
}

// NewEditorWidget returns a stub editor widget.
func NewEditorWidget() shell.Widget {
	return &EditorWidget{}
}

// Init satisfies shell.Widget.
func (w *EditorWidget) Init() tea.Cmd { return nil }

// Update satisfies shell.Widget.
func (w *EditorWidget) Update(msg tea.Msg) (shell.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
	case shell.OpenFileMsg:
		w.openedPath = msg.Path
	case shell.PTYInputMsg:
		w.lastKey = string(msg.Bytes)
	}
	return w, nil
}

// WriteKey satisfies shell.RawKeyWidget.
func (w *EditorWidget) WriteKey(msg tea.KeyMsg) tea.Cmd {
	w.lastKey = msg.String()
	return nil
}

// View satisfies shell.Widget.
func (w *EditorWidget) View() string {
	lines := []string{
		styles.Title.Render("Helix"),
		"",
		styles.Muted.Render("TODO: wire the PTY-backed editor pane."),
	}
	if w.openedPath != "" {
		lines = append(lines, "", "open file request: "+w.openedPath)
	}
	if w.lastKey != "" {
		lines = append(lines, "last raw key: "+w.lastKey)
	}
	return lipgloss.NewStyle().Padding(1).Render(strings.Join(lines, "\n"))
}

// FilesWidget is a stub for the file list shown beside Helix.
type FilesWidget struct {
	width   int
	height  int
	cursor  int
	entries []string
}

// NewFilesWidget returns a stub left-hand file widget.
func NewFilesWidget() shell.Widget {
	return &FilesWidget{
		entries: []string{
			"main.go",
			"internal/app/app.go",
			"tabs/helix/helix.go",
		},
	}
}

// Init satisfies shell.Widget.
func (w *FilesWidget) Init() tea.Cmd { return nil }

// Update satisfies shell.Widget.
func (w *FilesWidget) Update(msg tea.Msg) (shell.Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}
		case "down", "j":
			if w.cursor < len(w.entries)-1 {
				w.cursor++
			}
		case "enter":
			path := w.entries[w.cursor]
			return w, func() tea.Msg { return shell.OpenFileMsg{Path: path} }
		}
	}
	return w, nil
}

// View satisfies shell.Widget.
func (w *FilesWidget) View() string {
	lines := []string{
		styles.Title.Render("Files"),
		"",
		styles.Muted.Render("TODO: replace with the real file browser."),
		"",
	}
	for i, entry := range w.entries {
		prefix := "  "
		style := styles.Normal
		if i == w.cursor {
			prefix = "→ "
			style = styles.Selected
		}
		lines = append(lines, prefix+style.Render(entry))
	}
	lines = append(lines, "", styles.Help.Render("↑/↓ move • enter open • alt+→ helix"))
	return lipgloss.NewStyle().Padding(1).Render(strings.Join(lines, "\n"))
}
