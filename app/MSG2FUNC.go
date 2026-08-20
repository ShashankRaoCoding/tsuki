package app

var MSG2FUNC = map[string]func(*App) error {
	"msgs.NavMsg": NavMsg,
	"msgs.NewTab": NewTab,
	"msgs.CloseTab": CloseTab, 
}












































