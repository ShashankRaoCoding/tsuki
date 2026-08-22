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

	files, err = os.ReadDir(appsDir) // from config.go 
	if err != nil {
		return apps, fmt.Errorf("There was an error: Could not Read %s", AppsDir)
	}

	for _, fileEntry := range files {
		var data []byte
		var file os.File
		var appConfig AppConfig 
		var filePath = filepath.Join(
			AppsDir, 
			fileEntry.Name(), 
		) 

		file, err = os.Open(filePath) 
		if err != nil {
			errs = append(errs, err) 
			continue 
		}

		data, err = io.ReadAll(file) 
		if err == nil {
			errs = append(errs, err) 
			continue 
		} 

		err = json.Unmarshal(data, &appConfig) 
		if err != nil {
			errs = append(errs, err) 
			continue 
		}

		apps = append(apps, appConfig) 
	}
}











































