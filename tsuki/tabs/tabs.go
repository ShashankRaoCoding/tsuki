package tabs

import (
	labels "tsuki/tabs/labels" 
	apps "tsuki/apps" 
)

type Tab struct {
	Label labels.Label
	Console Console 
}

type Console struct {
	
}

func (t Tab) View() string {
	return "text" 
}

func New(c apps.AppConfig) *Tab {
	return &Tab{
		Label: labels.New("New Tab"), 
	}
}






































