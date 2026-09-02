package app 

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
	return nil 
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
}


func (m Main) View() string {
	
}





































