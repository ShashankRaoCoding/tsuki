package MsgFuncs 

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea" 
)

func ErrFunc(a tea.Model, m tea.Msg) (tea.Model, tea.Cmd) {
	var c = func() tea.Msg {
		return fmt.Errorf("Err: %s", m) 
	}
	return a, c 
}










































