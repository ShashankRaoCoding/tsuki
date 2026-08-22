package config

import (
	"github.com/ShashankRaoCoding/tsuki/shared/utils"
)

type Config struct {
	AppConfigDir string 
	
}

var ConfigPath = "config.json"
var CONFIG Config 

func init() {
	config, err := utils.ReadJSONToStruct(
		ConfigPath,
		func() Config {
			return Config{} 
		}, 
	)
}












































