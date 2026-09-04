package utils 

import (
	"fmt" 
	"os"
	"io"
	"encoding/json" 
	"io/fs" 
	"path/filepath" 
)

func ReadDirToStructs(filePath string, allS *[]any, f func() any) error {
	var files []fs.DirEntry
	var err error 
	var errs []error 
	var s any 

	files, err = os.ReadDir(filePath) 
	if err != nil {
		return err 
	}

	for _, file := range files {
		fullPath := filepath.Join(filePath, file.Name()) 
		s = f() 
		err = ReadJSONToStruct(fullPath, s) 
		if err != nil {
			errs = append(errs, err) 
			continue 
		}
		allS = append(allS, s) 
	}

	err = ErrsToErr(errs) 
	return err
}

func ReadJSONToStruct(filePath string, s any) error {
	var file *os.File
	var fileData []byte 
	var err error 

	file, err = os.Open(filePath) 
	if err != nil {
		return err
	}

	fileData, err = io.ReadAll(file) 
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, s)
	return err
}










































