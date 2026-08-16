package app

import (
	tab tsuki/tab
	msgs tsuki/msgs 
	appConfig tsuki/shared/appConfig 
)

type App struct {
	Tabs []tab.Tab
	Focus int // the tab index that is focussed 
	Apps map[string]appConfig.AppConfig
	Msgs chan msgs.Msg
}

func New() {
	apps, err := LoadApps() 
	launcher, err := tab.New(apps['tsuki-launcher']) // tsuki-launcher is an external dependency and not in scope for tsuki 
	a := App{
		Tabs: []tab.Tab{launcher}, 
		Focus: 0, 
		Apps: apps,
		Msgs: make(chan msgs.Msg, 1000) 
	} 
}

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

func (a *App) Listen() {
	// while true {
	//  i := input
	//     if i is ctrl + q, w, or t,
	//       create msgs.Cmd
	//     else,
	//       create msgs.Key
	//  a.Msgs <- msg 
	// }
}

func (a *App) 
































