package main

import (
	"embed"
	"fmt"
	"github.com/ShashankRaoCoding/kumo"
	"github.com/ShashankRaoCoding/tsuki/globals"
	"log"
	_ "github.com/ShashankRaoCoding/tsuki/handlerfuncs"
	"path/filepath"
)

//go:embed static
var FS embed.FS

func init() {
	globals.FS = FS
}

func main() {
	var err error

	// create kumo
	dbPath := fmt.Sprintf("%s/.tsuki.db", filepath.Dir(globals.BINARYPATH))
	s, err := kumo.New(
		kumo.WithDBPath(dbPath),
	)
	if err != nil {
		log.Println(err.Error())
	}

	// initialise database
	err = initialiseDBSchema(s.DB)
	if err != nil {
		log.Println(err.Error())
	}

	// start kumo
	log.Println(s.Listener.Addr())

	s.Paths = globals.PATHS

	s.Start()

}
