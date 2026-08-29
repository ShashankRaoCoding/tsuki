package app

import (
	"github.com/ShashankRaoCoding/tsuki/tabs" 
)

func renderTabLabels(a App) []string {
	logo := " Tsuki" 
	labels  := []string{} 
	labelWidth := (a.Width - 2 - len(logo)  - (len(" | ") * len(a.Tabs))) / len(a.Tabs)
	for _, tab := range a.Tabs {
		label := tab.Label.View(labelWidth) 
		labels = append(labels, label) 
	}

	return labels 
}










































