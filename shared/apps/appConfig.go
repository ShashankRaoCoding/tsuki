package appsConfig

import (
	"fmt" 
	"os" 
	"path/filepath" 
)

type AppConfig map[string]string

func LoadApps(appsDir string) ([]AppConfig, error) {
	var files []fs.DirEntry
	var err error
	var apps []AppConfig 
	var errs []error 

	files, err = os.ReadDir(appsDir) 
	if err != nil {
		return apps, fmt.Errorf("There was an error: Could not Read %s", AppsDir)
	}

	for _, fileEntry := range files {
		var appConfig AppConfig 
		var filePath = filepath.Join(
			AppsDir, 
			fileEntry.Name(), 
		) 

		apps = append(apps, appConfig) 
	}
}











































