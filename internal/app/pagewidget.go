package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/ShashankRaoCoding/tsuki/internal/page"
	"github.com/ShashankRaoCoding/tsuki/internal/shell"
)

// pageWidget adapts a page.Page so it can be used as a shell.Widget inside a Tab.
type pageWidget struct {
	p page.Page
}

func newPageWidget(p page.Page) *pageWidget {
	return &pageWidget{p: p}
}

// Init satisfies shell.Widget.
func (w *pageWidget) Init() tea.Cmd {
	return w.p.Init()
}

// Update satisfies shell.Widget.
func (w *pageWidget) Update(msg tea.Msg) (shell.Widget, tea.Cmd) {
	var cmd tea.Cmd
	if sz, ok := msg.(tea.WindowSizeMsg); ok {
		w.p.SetSize(sz.Width, sz.Height)
	}
	w.p, cmd = w.p.Update(msg)
	return w, cmd
}

// View satisfies shell.Widget.
func (w *pageWidget) View() string {
	return w.p.View()
}
