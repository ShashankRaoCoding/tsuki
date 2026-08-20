package main

import (
	"tsuki/app"
)

func main() {
	a := app.New() 
	err := a.Start() 
	if err != nil {
		panic(err) 
	}
}











































