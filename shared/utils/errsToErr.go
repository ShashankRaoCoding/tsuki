package utils

import (
	"strings" 
	"fmt" 
)

func ErrsToErr(errs []error) error {
	var _errs = []string{} 
	for _, err := range errs {
		if err != nil {
			_errs = append(_errs, err.Err()) 
		}
	}

	if len(_errs) == 0 {
		return nil 
	} else {
		return strings.Join(_errs, "\n")
	}
}











































