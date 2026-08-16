package app

import (
	// bubble tea related 
	tab tsuki/tab/tab
	msgs tsuki/app/msgs msgs 
)

type App struct {
	tabs []tab.Tab
	focus int // the tab index that is focussed 
	clis []map[string]string
}

func (a App) Init() {
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
	focus := a.tabs[a.focus].View() 
	return strings.Join([]string{tabs, focus}, "\n") 
}




































