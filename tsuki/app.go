package mainApp

import (
	"log"
	tea "github.com/charmbracelet/bubbletea" 
	msgs "tsuki/msgs"
	tabs "tsuki/tabs" 
)

type Main struct {
	Tabs []*tabs.Tab,
	Height int
	Width int
	Focus int 
	
}

func New() Main {
	m := Main{}
	m.Tabs = append(
		m.Tabs,
		tabs.New(
			apps.Apps["New Tab"], 
		), 
	)

	return m 
}

func (m Main) Init() tea.Cmd {
	log.Println("Init") 
	return nil 
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	f, o := msgs.Msg2Func[fmt.Sprintf("%t", msg)) ]
	if o == false {
		f, _ = msgs.Msg2Func["msgs.ErrMsg"] 
	}

	_m, c := f(m, msg)
	m, _ = _m.(Main) 
	return m, c
}


func (m Main) View() string {
	tabLabels := renderLabels(m) 
	content := m.Tabs[m.Focus].View()
	return tabLabels + "\n" + content 
}





































