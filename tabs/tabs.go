package tabs

import (
	labels "github.com/ShashankRaoCoding/tsuki/tabs/labels" 
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








































