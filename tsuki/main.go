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
	}
}










































