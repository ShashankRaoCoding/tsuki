package msgs

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea" 
	// msgs "githu.com/ShashankRaoCoding/tsuki/msgs"
	"log" 
)

func init() {
	Msg2Func["msgs.ErrMsg"] = ErrFunc
}

func ErrFunc(a tea.Model, m tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Error: %s", m) 
	return a, nil 
}










































