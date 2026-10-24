package globals

import (
	"os"
)

var BINARYPATH string

func init() {
	var err error

	BINARYPATH, err = os.Executable()
	if err != nil {
		panic(err)
	}
}
