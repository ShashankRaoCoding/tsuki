package apps

import (
	// "fmt"
	"os" 
	"log" 
	utils "tsuki/utils"
)

var APPS map[string]AppConfig 

type AppConfig struct {
	Name string
	Syntax string
	About string 
}

func init() {
	apps, err := LoadApps(".tsuki/apps/") 
	if err == nil {
		for _, c := range apps {
			APPS[c.Name] = c 
		}
		return 
	} else {
		log.Printf("Error: %s\n", err) 
		os.Exit(1) 
	}
}

func LoadApps(filePath string) ([]AppConfig, error) {
	_appConfigs := []any{} 
	appConfigs := []AppConfig {} 
	var err error 

	err = utils.ReadDirToStructs(
		filePath,
		&_appConfigs, 
		func() any {
			return AppConfig{} 
		}, 
	) 

	for _, _c := range _appConfigs {
		c, _ := _c.(AppConfig) 
		appConfigs = append(appConfigs, c) 
	}

	return appConfigs , err 
}












































