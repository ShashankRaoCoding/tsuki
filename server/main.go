package main

import (
	"github.com/ShashankRaoCoding/kumo/structs/server"
	"log"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Println(err.Error()) 
	}

	log.Println(s.Listener.Addr()) 
	s.Start() 
}










































