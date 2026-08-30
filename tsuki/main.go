package main

import (
	"log" 
	"fmt"
	"os"
	"github.com/ShashankRaoCoding/tsuki/app"
	tea "github.com/charmbracelet/bubbletea"
	// _ "github.com/ShashankRaoCoding/tsuki/msgs/msgFuncs"
)

func main() {
	f, err := os.Create(".tuski/errors") 
	if err == nil {
			log.SetOutput(f) 
	}
	defer f.Close() 

	_, err := tea.NewProgram(
		app.New(),
		tea.WithAltScreen(), 
	).Run() 

	log.Println(err.Error()) 
	os.Exit(0) 
}










































