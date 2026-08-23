package utils

import (
	"strings" 
	"fmt" 
)

func ErrsToErr(errs []error) error {
	var i = 0 
	for i < len(errs) {
		if errs[i] == nil {
			errs = append(errs[0:i], errs[1+i:]) 
			i = -1 + i 
		}
	}
}











































