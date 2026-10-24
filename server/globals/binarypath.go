package variables 

import (
	"os"
	"path/filepath"
)

var BINARYPATH string

func init() {
	BINARYPATH, err := os.Executable()
	if err != nil {
		panic(err)
	}
}







































