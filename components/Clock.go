package components

import (
	// "fmt" 
	"time" 
	"tsuki/globals" 
	"strings" 
)

type Clock struct {
	Unit time.Duration
	Progress int
	Total int
}

func (d *Clock) Start() {

	// calcuate seconds since start of the day 
    now := time.Now()

    // Get start of today (midnight)
    midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

    // Difference in seconds
    secondsToday := int(now.Sub(midnight).Seconds())

	d.Progress = secondsToday 
	for true {
		time.Sleep(d.Unit) 
		d.Progress++ 
	}
}

func (d *Clock) Update(event globals.Event) {
	// since day view doesn't expect inputs, this can be skipped 
	// pass 
	
}

func (d *Clock) View(width, height int) string {
	var output []string
	progress := int(float64(d.Progress)/float64(d.Total ) * float64(width) ) 
	for range progress {
		output = append(output, "»") 
	}
	for range (width-progress) {
		output = append(output, " ") 
	}
	return strings.Join(output, "") 
}





































