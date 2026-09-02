package tabs

import (
	labels "orca/tabs/labels" 
)

type Tab struct {
	Label labels.Label
	Console Console 
}

type Console {
	
}

func (c Console) View() string {
	return "text" 
}

func New() Tab {
	return Tab{
		Label: labels.New("New Tab"), 
	}
}






































