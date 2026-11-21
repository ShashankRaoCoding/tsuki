package handlerfuncs

import (
	"github.com/ShashankRaoCoding/kumo"
	"github.com/ShashankRaoCoding/tsuki/globals"
	"net/http"
)

func init() {
	
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// serve the index.html file
	kumo.ServeFromEmbed(w, r, globals.FS, "static/index.html", nil)
}
