package app

import (
	// bubble tea related 
	tab tsuki/tsuki/tab
	msgs tsuki/tsuki/msgs 
)

// func New() Main {
	// creates a Main App 
	// m := Main{} 
	// reads apps/* and constructs Main.Apps 
	// return m 
// }

type Main struct {
	Tabs map[int]tab.Tab
	Focus int // the tab index that is focussed 
	Apps map[string]appConfig
}

type appConfig map[string]string

func (m *Main) Start() {
	// reads CLIs/* and adds the CLIs 
	// starts tab 'launcher' 
}

func (m *Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// if reserved, interpret
	// else forward to focus tab with
	// a.Tabs[a.Focus] = a.Tabs[a.Focus].Update(tea.Msg)
	// return a,  
	// or better equivalent 
}


func (m *Main) View() string {
	tabs := a.RenderTabs() 
	focus := a.Tabs[a.Focus].View() 
	return strings.Join([]string{tabs, focus}, "\n") 
}




































