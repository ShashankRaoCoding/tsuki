package main

import (
	"fmt"
	"os"
	"github.com/ShashankRaoCoding/tsuki/app"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var err error
	_, err = tea.NewProgram(
		app.New(),
		tea.WithAltScreen, 
	)

	app.New().Init() 
	app.New().Update()
	app.New.View()
	
	if err == nil {
		os.Exit(0) 
	}

	fmt.Printf(err)
	os.Exit(1) 
}










































