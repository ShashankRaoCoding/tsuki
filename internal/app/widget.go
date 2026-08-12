package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

// Widget is the shared pane interface used by tabs.
type Widget interface {
	Init() tea.Cmd
	Update(tea.Msg) (Widget, tea.Cmd)
	View() string
}

// RawKeyWidget writes key input directly instead of consuming tea.KeyMsg in Update.
type RawKeyWidget interface {
	WriteKey(tea.KeyMsg) tea.Cmd
}

// FocusZone describes which application area currently owns keyboard focus.
type FocusZone int

const (
	FocusGlobal FocusZone = iota
	FocusTabLeft
	FocusTabRight
)

// WidgetTarget identifies which pane should receive a routed message.
type WidgetTarget int

const (
	WidgetTargetNone WidgetTarget = iota
	WidgetTargetLeft
	WidgetTargetRight
)

// OpenFileMsg requests that the main editor widget open a path.
type OpenFileMsg struct {
	Path string
}

// PTYInputMsg models bytes destined for a PTY-backed widget.
type PTYInputMsg struct {
	Bytes []byte
}

// SetTabTitleMsg allows widgets to request a title change via the parent tab.
type SetTabTitleMsg struct {
	Title string
}

// FocusTargetForKey is the pure focus-routing function used by tabs.
func FocusTargetForKey(zone FocusZone, _ tea.KeyMsg, hasLeft bool) WidgetTarget {
	switch zone {
	case FocusTabLeft:
		if hasLeft {
			return WidgetTargetLeft
		}
		return WidgetTargetRight
	case FocusTabRight:
		return WidgetTargetRight
	default:
		return WidgetTargetNone
	}
}

// HelixWidget is a stub for a future PTY-backed editor pane.
type HelixWidget struct {
	width      int
	height     int
	openedPath string
	lastKey    string
}

// NewHelixWidget returns a stub editor widget.
func NewHelixWidget() Widget {
	return &HelixWidget{}
}

// Init satisfies Widget.
func (w *HelixWidget) Init() tea.Cmd { return nil }

// Update satisfies Widget.
func (w *HelixWidget) Update(msg tea.Msg) (Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
	case OpenFileMsg:
		w.openedPath = msg.Path
	case PTYInputMsg:
		w.lastKey = string(msg.Bytes)
	}
	return w, nil
}

// WriteKey satisfies RawKeyWidget.
func (w *HelixWidget) WriteKey(msg tea.KeyMsg) tea.Cmd {
	w.lastKey = msg.String()
	return nil
}

// View satisfies Widget.
func (w *HelixWidget) View() string {
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

// CWDWidget is a stub for a future cwd/tree navigation pane.
type CWDWidget struct {
	width   int
	height  int
	cursor  int
	entries []string
}

// NewCWDWidget returns a stub left-hand navigation widget.
func NewCWDWidget() Widget {
	return &CWDWidget{
		entries: []string{
			"main.go",
			"internal/app/app.go",
			"TODO: cwd tree",
		},
	}
}

// Init satisfies Widget.
func (w *CWDWidget) Init() tea.Cmd { return nil }

// Update satisfies Widget.
func (w *CWDWidget) Update(msg tea.Msg) (Widget, tea.Cmd) {
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
			return w, func() tea.Msg { return OpenFileMsg{Path: path} }
		}
	}
	return w, nil
}

// View satisfies Widget.
func (w *CWDWidget) View() string {
	lines := []string{
		styles.Title.Render("CWD"),
		"",
		styles.Muted.Render("TODO: replace with a real cwd tree widget."),
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
	lines = append(lines, "", styles.Help.Render("↑/↓ move • enter open • alt+→ editor"))
	return lipgloss.NewStyle().Padding(1).Render(strings.Join(lines, "\n"))
}
