package shared 

import (
	"fmt" 
	"json/Encoding" 
	"os" 
	"io" 
	"io/fs" 
	"path/filepath" 
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

	for _, fileEntry := range files {
		var data []byte
		var file os.File

		file, err = os.Open(filepath.Join(
			AppsDir, 
			fileEntry.Name(), ) 
		) 
		data, err = io.ReadAll(file) 
		if err == nil {
			
		} else {
			errs = append(
				errs,
				fmt.Errorf("Could not read %s", file) , 
			)
		}
	}
}











































