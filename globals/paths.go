package globals

import "net/http"

var PATHS = map[string]func(w http.ResponseWriter, r *http.Request){}
