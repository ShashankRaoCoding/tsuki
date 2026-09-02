package appsConfig

import (
	"fmt"
	"os" 
	utils "tsuki/shared/utils"
)

var Apps map[string]AppConfig 

type AppConfig map[string]string 

func LoadApps(filePath string) ([]AppConfig, error) {
	var appConfig  = AppConfig {} 
	var appConfigs []AppConfig 
	var err error 

	appConfigs, err = utils.ReadDirToStructs(
		filePath,
		func() AppConfig{
			return AppConfig{} 
		}, 
	) 

}












































