package utils 

import (
	"fmt" 
	"os"
	"io"
	"encoding/json" 
)

func ReadJSONToStruct(filePath, structType) (any, err ) {
	var file os.File
	var fileData []byte 
	var err error 
	var s = make(structType) 

	file, err = os.Open(filePath) 
	if err != nil {
		return s, err 
	}

	fileData, err = io.ReadAll(file) 
	if err != nil {
		return s, err 
	}

	err = json.Unmarshal(fileDtaa, &s)
	return s, err 
}










































