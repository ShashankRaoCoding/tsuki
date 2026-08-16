package app

import (
	tab tsuki/tab
	msgs tsuki/msgs 
	appConfig tsuki/shared/appConfig 
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
	Apps map[string]appConfig.AppConfig
	Msgs chan msgs.Msg
}

// type appConfig map[string]string

func New() {
	apps, err := LoadApps() 
	launcher, err := tab.New(apps['tsuki-launcher'])
	a := App{
		Tabs: []tab.Tab{launcher}, 
		Focus: 0, 
		Apps: apps,
		Msgs: make(chan msgs.Msg, 1000) 
	} 
	// Create msgs chan 
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




































