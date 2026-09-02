package appsConfig

import (
	"fmt"
	"os" 
	utils "tsuki/shared/utils"
)

var Apps map[string]AppConfig 

type AppConfig map[string]string 

func init() {
	
}

func LoadApps(filePath string) ([]AppConfig, error) {
	_appConfigs := []any{} 
	appConfigs := []AppConfig {} 
	var err error 

	err = utils.ReadDirToStructs(
		filePath,
		func() AppConfig{
			return AppConfig{} 
		}, 
	) 

}












































