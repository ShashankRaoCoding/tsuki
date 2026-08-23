package appsConfig

import (
	"fmt"
	"os" 
	utils "github.com/ShashankRaoCoding/tsuki/shared/utils"
)

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












































