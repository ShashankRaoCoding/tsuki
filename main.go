package main

import (
	"fmt"
	"github.com/ShashankRaoCoding/kumo/config"
	"github.com/ShashankRaoCoding/kumo/server"
	"github.com/ShashankRaoCoding/tsuki/globals"
	"log"
	"path/filepath"
)

func main() {
	var err error

	// create server
	dbPath := fmt.Sprintf("%s/.tsuki.db", filepath.Dir(globals.BINARYPATH))
	s, err := server.New(
		config.WithDBPath(dbPath),
	)
	if err != nil {
		log.Println(err.Error())
	}

	// initialise database
	err = initialiseDBSchema(s.DB)
	if err != nil {
		log.Println(err.Error())
	}

	// start server
	log.Println(s.Listener.Addr())
	s.Start()

}
