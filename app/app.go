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

	f, o = msgs.Msg2Func[m]
	if o == false {
		f, _ = msgs.Msg2Func["Err"] 
	}

	a, c = f(a, m) 
	return a, c
}

func (a App) View() string {
	return a.Render() 
}

func (a App) Render() string {
	return lipgloss.JoinVertical(
		tabs,
		content, 
	)
}


































