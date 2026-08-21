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

	fmt.Printf(err)
	os.Exit(0) 
}










































