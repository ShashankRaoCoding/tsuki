package msgs

import (
	"fmt"
	msgFuncs "github.com/ShashankRaoCoding"
	tea "github.com/charmbracelet/tea"
)

type MsgFunc func(tea.Model, tea.Msg) (tea.Model, tea.Cmd) 

var Msg2Func = map[string]MsgFunc{
	
}








































