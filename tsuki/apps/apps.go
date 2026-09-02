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
	appConfig  := []AppConfig {} 
	appConfigs := []AppConfig {} 
	var err error 

	appConfigs, err = utils.ReadDirToStructs(
		filePath,
		func() AppConfig{
			return AppConfig{} 
		}, 
	) 

}












































