package main

import (
	"github.com/ShashankRaoCoding/kumo/server"
	"github.com/ShashankRaoCoding/kumo/config"
	"github.com/ShashankRaoCoding/tsuki-server/funcs/initialise"
	"fmt" 
	"log"
	"path/filepath"
	"github.com/ShashankRaoCoding/tsuki/utils/variables"
)

func main() {
	var err error 

	// create server 
	dbPath = fmt.Sprintf("%s/.tsuki.db", filepath.Dir(variables.BINARYPATH))
	s, err := server.New(
		config.WithDBPath(dbPath),
	)
	if err != nil {
		log.Println(err.Error()) 
	}
	
	// initialise database  
	err = initialise.DBSchema(s.DB)
	if err != nil {
		log.Println(err.Error())
	}

	// start server 
	log.Println(s.Listener.Addr()) 
	s.Start() 

}








































