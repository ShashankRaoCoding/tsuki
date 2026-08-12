package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ShashankRaoCoding/tsuki/internal/shell"
	"github.com/ShashankRaoCoding/tsuki/internal/styles"
)

type appOption struct {
	id          string
	title       string
	description string
}

type launcherWidget struct {
	options []appOption
	cursor  int
	width   int
	height  int
}

func newLauncherWidget(options []appOption) shell.Widget {
	return &launcherWidget{options: append([]appOption(nil), options...)}
}

func (w *launcherWidget) Init() tea.Cmd { return nil }

func (w *launcherWidget) Update(msg tea.Msg) (shell.Widget, tea.Cmd) {
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
			if w.cursor < len(w.options)-1 {
				w.cursor++
			}
		case "enter", " ":
			if len(w.options) == 0 {
				return w, nil
			}
			option := w.options[w.cursor]
			return w, func() tea.Msg { return shell.StartAppMsg{AppID: option.id} }
		}
	}
	return w, nil
}

func (w *launcherWidget) View() string {
	lines := []string{
		styles.Title.Render("Start an app"),
		"",
		styles.Muted.Render("Open a tab target from /tabs."),
		"",
	}

	for i, option := range w.options {
		prefix := "  "
		titleStyle := styles.Normal
		descStyle := styles.Muted
		if i == w.cursor {
			prefix = "→ "
			titleStyle = styles.Selected
			descStyle = styles.Subtitle
		}

		lines = append(lines, prefix+titleStyle.Render(option.title))
		lines = append(lines, "  "+descStyle.Render(option.description))
		lines = append(lines, "")
	}

	lines = append(lines, styles.Help.Render("↑/↓ move • enter start • ctrl+t new tab • ctrl+w close"))
	content := lipgloss.NewStyle().Padding(1, 2).Render(strings.Join(lines, "\n"))

	if w.width > 0 && w.height > 0 {
		return lipgloss.Place(w.width, w.height, lipgloss.Center, lipgloss.Center, content)
	}
	return content
}
