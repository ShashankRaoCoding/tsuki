package utils 

import (
	"fmt" 
	"os"
	"io"
	"encoding/json" 
	"io/fs" 
)

func ReadDirToStructs(filePath string, s any) (any, err) {
	var files []fs.DirEntry
	var err error 
	var errs []error 

	files, err = os.ReadDir(filePath) 
	for _, file := files 
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










































