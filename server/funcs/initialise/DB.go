package initialise 

import (
	"log" 
	"github.com/ShashankRaoCoding/kumo/structs/server"
	"github.com/ShashankRaoCoding/tsuki/utils/db"

)

// this file initialises databases required by tsuki 

func DB(s *server.Server) error {
	var err error 

	// create a new sqlite db if not exists else open existing at .USERDATA.db 
	s.DB, err = sqlite.Open(s.Config.DBPath)
	if err != nil {
		log.Println(err.Error())
	}

	// initialise schema for the database
	err = 
	if err != nil {
		log.Println(err.Error())
	}

	return err 
}






































