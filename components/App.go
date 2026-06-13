package components

import (
	
	"log"
	// "time" 
	"strings"

	"tsuki/globals"
	"github.com/gdamore/tcell/v2"
)

type Component interface {
	Start() 
	Update(e globals.Event)
	View(width, height int) string
}

// struct holds app structure
type App []Component
var SCREEN tcell.Screen 
var FOCUS Component 
var RUN bool = true 
var keyNames = map[tcell.Key]string{
	tcell.KeyEscape: "Escape",
	tcell.KeyEnter: "Enter",
	tcell.KeyUp: "Up",
	tcell.KeyDown: "Down",
	tcell.KeyLeft: "Left",
	tcell.KeyRight: "Right",
	tcell.KeyCtrlC: "Ctrl+C",
}

// start app 
func (a *App) Start() {

	// screen to handle event polling 
	var err error
	SCREEN, err = tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	// Initialise the screen 
	if err := SCREEN.Init(); err != nil {
		log.Fatal(err)
	}
	defer SCREEN.Fini()

	// start all components
	for _, c := range *a {
		go c.Start() 
	}

	// start async event listening 
	// eventsChannel := make(chan string, 100)
	go a.Listen()

	// start the rendering 
	var output string 
	var height int
	var width int

	for SCREEN == nil {
		// pass wait for screen to initialise 
	}
		
	width, height = SCREEN.Size() 

	for event := range globals.EventsChannel {
		go FOCUS.Update(event) 
		output = a.View(width, height) // render 
		a.Print(output, width) 
		// time.Sleep(time.Second) 
	} 

	return 
}

func (a *App) View(width, height int) string {
	var output []string 
	
	for _, c := range *a {
		output = append(output, c.View(width, height)) 
	}

	return strings.Join(output, "\n") 
}


func (a *App) Listen() {

	
	for {
		// loops till Escape 
		event := SCREEN.PollEvent()
		if ev, ok := event.(*tcell.EventKey); ok {

			// keypress, decipher to string 
			keyName, ok := keyNames[ev.Key()]
			if ok == false {
				keyName = string(ev.Rune())
			}

			// RUN = false 
			if keyName == "Escape" {
				RUN = false 
				return
			}

			// regular key press, redirect to focused 
			// FOCUS.Update(keyName) 
			globals.EventsChannel <- globals.Event{
				Type: "key",
				Content: keyName,
			} 

			
		}
	}
}

func (a *App) Print(text string, width int) {
	x := 0
	y := 0 
	style := tcell.StyleDefault
	col := x
	row := y

	for _, r := range text {
		if col >= x+width {
			row++
			col = x
		}
		SCREEN.SetContent(col, row, r, nil, style)
		col++
	}

	SCREEN.Show()
}



































