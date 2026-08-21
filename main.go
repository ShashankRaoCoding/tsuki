package main

import (
	"os"
	"fmt" 
	"github.com/ShashankRaoCoding/tuski/app" 
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := tea.NewProgram(
		app.New(),
		tea.WithAltScreen, 
	)

	if err == nil {
		os.Exit(0) 
	} else {
		fmt.Printf(err)
		os.Exit(1) 
	}
}










































