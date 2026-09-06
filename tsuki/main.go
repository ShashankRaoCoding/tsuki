package main

import (
	tea "github.com/charmbracelet/bubbletea" 
	"os"
	"fmt"
	"log" 
)

func main() {
	f, err := os.Create(".tsuki/logs.txt") 
	if err == nil {
		log.SetOutput(f) 
	} else {
		fmt.Println("Error: %s", err, ", printing logs to stdout") 
	}

	t := tea.NewProgram(
		New(),
		tea.WithAltScreen(), 
	)

	_, err = t.Run() 
}










































