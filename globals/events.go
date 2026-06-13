package globals

var EventsChannel = make(chan Event, 100) 

type Event struct {
	Type string
	Content string
}







































