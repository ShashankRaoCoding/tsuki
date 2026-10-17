package handlerfuncs 

import (
	"net/http"
	"github.com/ShashankRaoCoding/tsuki/globals/variables"
)

func init() {
	variables.PATHS["/"] = indexHandler
	variables.PATHS["/index"] = indexHandler
	variables.PATHS["/index.html"] = indexHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// serve the index.html file 
	http.ServeFile(w, r, "static/index.html") 
}





































