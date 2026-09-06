package main

import (
	"fmt" 
	"log"
	tea "github.com/charmbracelet/bubbletea" 
	msgs "tsuki/msgs"
	tabs "tsuki/tabs" 
	utils "tsuki/utils" 
	apps "tsuki/apps" 
)

type Main struct {
	Tabs []*tabs.Tab
	Height int
	Width int
	Focus int 
	
}

type Config struct {
	AppsDir string 
	
}

var ConfigPath = "config.json"
var CONFIG Config 

func init() {
	err := utils.ReadJSONToStruct(
		ConfigPath,
		CONFIG, 
	)

	if err == nil {
		return 
	}

	panic(err) 
}

func New() Main {
	m := Main{}
	tab := tabs.New(apps.APPS["New Tab"])
	m.Tabs = append(
		m.Tabs,
		tab, 
	)

	return m 
}

func (m Main) Init() tea.Cmd {
	log.Println("Init") 
	return nil 
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	f, o := msgs.Msg2Func[fmt.Sprintf("%t", msg) ]
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


func renderLabels(a Main) string {
	logo := " tsuki" 
	labels  := "" 
	labelWidth := (a.Width - 2 - len(logo)  - (len(" | ") * len(a.Tabs))) / len(a.Tabs)
	for _, tab := range a.Tabs {
		label := tab.Label.Render(labelWidth) 
		labels = labels + label
	}

	return labels 
}



































