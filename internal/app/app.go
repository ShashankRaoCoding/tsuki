package app

import (
	// bubble tea related 
	tab tsuki/internal/tab
	msgs tsuki/internal/msgs 
)

type App struct {
	Tabs map[string]tab.Tab
	Focus string // the app name that is focussed 
	Apps []map[string]string
}

func (a App) Init() (tea.Model, tea.Cmd) {
	// reads CLIs/* and adds the CLIs 
	// starts tab 'launcher' 
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// if reserved, interpret
	// else forward to focus tab with
	// return a, a.tabs[a.focus].Update(tea.Msg) 
	// or better equivalent 
}


func (a App) View() string {
	tabs := a.RenderTabs() 
	focus := a.Tabs[a.Focus].View() 
	return strings.Join([]string{tabs, focus}, "\n") 
}




































