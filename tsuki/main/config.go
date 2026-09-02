package tsuki

import (
	"tsuki/utils"
)

type Config struct {
	AppsDir string 
	
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












































