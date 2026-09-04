package tabs

import (
	labels "tsuki/tabs/labels" 
)

type Tab struct {
	Label labels.Label
	Console Console 
}

type Console struct {
	
}

func (c Console) View() string {
	return "text" 
}

func New(c AppConfig) *Tab {
	return &Tab{
		Label: labels.New("New Tab"), 
	}
}






































