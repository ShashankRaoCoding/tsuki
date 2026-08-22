package shared 

import (
	"fmt" 
	"json/Encoding" 
	"os" 
	"io/fs" 
)

type AppConfig map[string]string

func Load() AppConfig {
	var files []fs.DirEntry
	var err error

	files, err = os.ReadDir(AppsDir) // from config.go 
	if err != nil {
		fmt.Printf("There was an error: Could not Read %s", AppsDir)
		os.Exit(1) 
	}
}











































