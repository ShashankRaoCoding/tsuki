package appsConfig

import (
	"fmt"
	"os" 
	utils "tsuki/utils"
)

var Apps map[string]AppConfig 

type AppConfig map[string]string 

func init() {
	Apps = LoadApps("tsukki/apps/") 
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












































