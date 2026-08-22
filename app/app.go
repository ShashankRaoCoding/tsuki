package app 

import (
	msgs "github.com/ShashankRaoCoding/tsuki/msgs"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	shared "github.com/ShashankRaoCoding/tsuki/shared"
)

type App struct {
	Height int 
	Widht int 
	Tabs []*tab.Tab
	Apps map[string]shared.AppConfig
}

func (a App) Init() tea.Msg {
	apps, err = shared.LoadApps() 

	return nil 
}

func (a App) Update(m tea.Msg) (tea.Model, tea.Cmd) {
	var f msgs.MsgFunc
	var o bool
	var c tea.Cmd 

	f, o = msgs.Msg2Func[m]
	if o == false {
		f, _ = msgs.Msg2Func["Err"] 
	}

	a, c = f(a, m) 
	return a, c
}

func (a App) View() string {
	return a.Render() 
}

func (a App) Render() string {
	tab := app.RenderTabs()
	content := app.Tabs[app.Focus.View()] 
	return lipgloss.JoinVertical(
		tabs,
		content, 
	)
}


































