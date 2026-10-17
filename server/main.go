package main

import (
	"github.com/ShashankRaoCoding/kumo/structs/server"
	"github.com/ShashankRaoCoding/kumo/funcs/initialise"
	"fmt" 
	"log"
	"path/filepath"
	"github.com/ShashankRaoCoding/tsuki/utils/variables"
)

func main() {
	var err error 

	s, err := server.New()
	if err != nil {
		log.Println(err.Error()) 
	}

	s.Config.DBPath = fmt.Sprintf("%s/.tsuki.db", filepath.Dir(variables.BINARYPATH))
	initialise.DB(s)

	log.Println(s.Listener.Addr()) 
	s.Start() 

}








































