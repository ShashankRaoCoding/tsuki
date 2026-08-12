package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const minSidebarWidth = 24

// Tab contains the optional left pane and required right pane for one shell tab.
type Tab struct {
	title      string
	left       Widget
	right      Widget
	width      int
	height     int
	leftWidth  int
	rightWidth int
}

func newDefaultTab(id int) Tab {
	return Tab{
		title: fmt.Sprintf("Tab %d", id),
		left:  NewCWDWidget(),
		right: NewHelixWidget(),
	}
}

// Init initialises both widgets owned by the tab.
func (t Tab) Init() tea.Cmd {
	cmds := []tea.Cmd{t.right.Init()}
	if t.left != nil {
		cmds = append(cmds, t.left.Init())
	}
	return tea.Batch(cmds...)
}

// Update handles focus changes, cross-pane messages, resize, and widget forwarding.
func (t Tab) Update(msg tea.Msg, zone FocusZone) (Tab, FocusZone, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmd := t.applySize(msg.Width, msg.Height)
		return t, normalizeFocusZone(zone, t), cmd
	case tea.KeyMsg:
		if nextZone, handled := t.toggleFocusZone(zone, msg); handled {
			return t, nextZone, nil
		}
		return t.forwardKey(msg, zone)
	case OpenFileMsg:
		return t.deliverToRight(msg, normalizeFocusZone(zone, t))
	case SetTabTitleMsg:
		t.title = msg.Title
		return t, normalizeFocusZone(zone, t), nil
	default:
		cmd := t.broadcast(msg)
		return t, normalizeFocusZone(zone, t), cmd
	}
}

// View renders the tab body as either a split view or a full-width right pane.
func (t Tab) View() string {
	if t.right == nil {
		return ""
	}

	right := t.renderPane(t.right.View(), t.rightWidth, t.height)
	if t.left == nil {
		return right
	}

	left := t.renderPane(t.left.View(), t.leftWidth, t.height)
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (t *Tab) applySize(width, height int) tea.Cmd {
	t.width = width
	t.height = height

	if t.left == nil {
		t.leftWidth = 0
		t.rightWidth = maxInt(width, 0)
		return t.resizeRight()
	}

	leftWidth := width / 3
	if leftWidth < minSidebarWidth {
		leftWidth = minSidebarWidth
	}
	if width > 1 && leftWidth >= width {
		leftWidth = width - 1
	}
	if leftWidth < 0 {
		leftWidth = 0
	}

	t.leftWidth = leftWidth
	t.rightWidth = width - leftWidth
	if t.rightWidth < 0 {
		t.rightWidth = 0
	}

	return tea.Batch(t.resizeLeft(), t.resizeRight())
}

func (t Tab) toggleFocusZone(zone FocusZone, msg tea.KeyMsg) (FocusZone, bool) {
	switch msg.String() {
	case "alt+left":
		if t.left != nil {
			return FocusTabLeft, true
		}
		return FocusTabRight, true
	case "alt+right":
		return FocusTabRight, true
	default:
		return zone, false
	}
}

func (t *Tab) forwardKey(msg tea.KeyMsg, zone FocusZone) (Tab, FocusZone, tea.Cmd) {
	target := FocusTargetForKey(zone, msg, t.left != nil)

	switch target {
	case WidgetTargetLeft:
		if t.left == nil {
			return *t, normalizeFocusZone(zone, *t), nil
		}
		if raw, ok := t.left.(RawKeyWidget); ok {
			return *t, normalizeFocusZone(zone, *t), raw.WriteKey(msg)
		}
		var cmd tea.Cmd
		t.left, cmd = t.left.Update(msg)
		return *t, normalizeFocusZone(zone, *t), cmd
	case WidgetTargetRight:
		if raw, ok := t.right.(RawKeyWidget); ok {
			return *t, normalizeFocusZone(zone, *t), raw.WriteKey(msg)
		}
		var cmd tea.Cmd
		t.right, cmd = t.right.Update(msg)
		return *t, normalizeFocusZone(zone, *t), cmd
	default:
		return *t, normalizeFocusZone(zone, *t), nil
	}
}

func (t *Tab) deliverToRight(msg tea.Msg, zone FocusZone) (Tab, FocusZone, tea.Cmd) {
	if t.right == nil {
		return *t, zone, nil
	}
	var cmd tea.Cmd
	t.right, cmd = t.right.Update(msg)
	return *t, zone, cmd
}

func (t *Tab) broadcast(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, 0, 2)
	if t.left != nil {
		var cmd tea.Cmd
		t.left, cmd = t.left.Update(msg)
		cmds = append(cmds, cmd)
	}
	if t.right != nil {
		var cmd tea.Cmd
		t.right, cmd = t.right.Update(msg)
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

func (t *Tab) resizeLeft() tea.Cmd {
	if t.left == nil {
		return nil
	}
	var cmd tea.Cmd
	t.left, cmd = t.left.Update(tea.WindowSizeMsg{
		Width:  t.leftWidth,
		Height: t.height,
	})
	return cmd
}

func (t *Tab) resizeRight() tea.Cmd {
	if t.right == nil {
		return nil
	}
	var cmd tea.Cmd
	t.right, cmd = t.right.Update(tea.WindowSizeMsg{
		Width:  t.rightWidth,
		Height: t.height,
	})
	return cmd
}

func (t Tab) renderPane(content string, width, height int) string {
	style := lipgloss.NewStyle()
	if width > 0 {
		style = style.Width(width).MaxWidth(width)
	}
	if height > 0 {
		style = style.Height(height).MaxHeight(height)
	}
	return style.Render(content)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
