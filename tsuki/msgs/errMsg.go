package msgFuncs

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea" 
	msgs "github.com/ShashankRaoCoding/tsuki/msgs"
	"log" 
)

func init() {
	msgs.Msg2Func["msgFuncs.ErrMsg"] = ErrFunc
}

func ErrFunc(a tea.Model, m tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("Error: %s", m) 
	return a, nil 
}










































