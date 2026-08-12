package shell

import tea "github.com/charmbracelet/bubbletea"

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

// StartAppMsg requests that the current launcher tab start a specific app.
type StartAppMsg struct {
	AppID string
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
