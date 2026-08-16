package app

import (
	// bubble tea related 
	tab tsuki/tab
	msgs tsuki/msgs 
	shared tsuki/shared 
)

// func New() App {
	// creates a App App 
	// m := App{} 
	// reads apps/* and constructs App.Apps 
	// return m 
// }

type App struct {
	Tabs []tab.Tab
	Focus int // the tab index that is focussed 
	Apps map[string]shared.AppConfig
	Msgs chan msgs.Msg
}

// type appConfig map[string]string

func (a *App) Start() {
	// reads CLIs/* and adds the CLIs 
	// starts tab 'launcher' 
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// if reserved, interpret
	// else forward to focus tab with
	// a.Tabs[a.Focus] = a.Tabs[a.Focus].Update(tea.Msg)
	// return a,  
	// or better equivalent 
}


func (a *App) View() string {
	tabs := a.RenderTabs() 
	focus := a.Tabs[a.Focus].View() 
	return strings.Join([]string{tabs, focus}, "\n") 
}




































