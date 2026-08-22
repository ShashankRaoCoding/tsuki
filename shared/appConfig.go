package shared 

import (
	"fmt" 
	"json/Encoding" 
	"os" 
	"io" 
	"io/fs" 
)

type AppConfig map[string]string

func LoadApps() ([]AppConfig, error) {
	var files []fs.DirEntry
	var err error
	var apps []AppConfig 
	var errs []error 

	files, err = os.ReadDir(AppsDir) // from config.go 
	if err != nil {
		return apps, fmt.Errorf("There was an error: Could not Read %s", AppsDir)
	}

	for _, file := range files {
		var data []byte
		data, err = io.ReadAll(file) 
		if err == nil {
			
		} else {
			
		}
	}
}











































