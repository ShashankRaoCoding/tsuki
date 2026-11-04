package handlerfuncs 

import (
	"net/http"
	"github.com/ShashankRaoCoding/tsuki/globals"
	"github.com/ShashankRaoCoding/kumo"
)

func init() {
	globals.PATHS["/"] = indexHandler
	globals.PATHS["/index"] = indexHandler
	globals.PATHS["/index.html"] = indexHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// serve the index.html file 
	kumo.ServeFromEmbed(w, r, globals.FS, "static/index.html") 
}





































