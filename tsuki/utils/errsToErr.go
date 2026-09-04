package utils

import (
	"strings" 
	"fmt" 
)

func ErrsToErr(errs []error) error {
	var _errs = []string{} 
	for _, err := range errs {
		if err != nil {
			_errs = append(_errs, err.Error()) 
		}
	}

	if len(_errs) == 0 {
		return nil 
	} else {
		return fmt.Errorf("%s", strings.Join(_errs, "\n")) 
	}
}











































