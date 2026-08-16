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

func New() {
	// a := App{} 
	// Focus: 0 
	// Create msgs chan 
	// reads apps/*.json to AppConfig structs and adds the AppConfig structs for each .json to a.Apps 
	// starts tab 'launcher' 
}

// func (a *App) Start() int {
	// go a.Listen() 
// }

func (a *App) Start() error {
	go a.Listen() // translate inputs to msgs 
	
	for msg := range a.Msgs {
	//   if msg is navMsg, change focus += msg.Delta 
	//   if msg is a command, interpret
	//   if msg is a key press forward to a.Tabs[a.Focus]
	}
}

func (a *App) View() string {
	tabs := a.RenderTabs() 
	focus := a.Tabs[a.Focus].View() 
	return strings.Join([]string{tabs, focus}, "\n") 
}




































