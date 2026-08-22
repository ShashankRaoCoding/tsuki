package utils 

import (
	"fmt" 
	"os"
	"io"
	"encoding/json" 
)

func ReadDirToStructs(filePath string, s any) (any, err) {
	
}

func ReadJSONToStruct(filePath string, s any) (any, err) {
	var file os.File
	var fileData []byte 
	var err error 

	file, err = os.Open(filePath) 
	if err != nil {
		return s, err
	}

	fileData, err = io.ReadAll(file) 
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(fileData, &s)
	return s, err
}










































