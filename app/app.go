package app 

import (
	msgs "github.com/ShashankRaoCoding/tsuki/msgs"
	tea "github.com/charmbracelet/bubbletea"

)

func (a App) Init() tea.Msg {
	return nil 
}

func (a App) Update(m tea.Msg) (tea.Model, tea.Cmd) {
	var f msgs.MsgFunc
	var o bool
	var c tea.Cmd 

	return a, c
}

func (a App) View() string {
	return a.Render() 
}




































