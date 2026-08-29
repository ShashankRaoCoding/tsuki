package app 

import (
	msgs "github.com/ShashankRaoCoding/tsuki/msgs"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	config "github.com/ShashankRaoCoding/tsuki/shared/config" 
	appsConfig "github.com/ShashankRaoCoding/tsuki/shared/appsConfig"
	tabs "github.com/ShashankRaoCoding/tsuki/tabs" 
	msgFuncs "github.com/ShashankRaoCoding/tsuki/msgs/msgFuncs" 
)

func New() App {
	return App{
		Tabs: []*tabs.Tab{
			tabs.New(), 
		}, 
	} 
}

type App struct {
	Height int 
	Width int 
	Tabs []*tab.Tab
	Apps map[string]appsConfig.AppConfig
}

func (a App) Init() tea.Cmd {
	_apps, err = appsConfig.LoadApps(".apps/") 
	for _, app := range _apps {
		a.Apps[app.Name] = app
	}

	return nil 
}

func (a App) Update(m tea.Msg) (tea.Model, tea.Cmd) {
	var f msgs.MsgFunc
	var o bool
	var c tea.Cmd 

	f, o = msgs.Msg2Func[m]
	if o == false {
		f, _ = msgs.Msg2Func["msgFuncs.ErrMsg"] 
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


































