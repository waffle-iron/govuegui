package govuegui

import (
	"net/http"

	"github.com/as27/golib/js/vuejsdev"
	"github.com/gorilla/mux"
)

// PathPrefix defines the prefix for the all gui specific tasks
var PathPrefix = "/govuegui"

// ServerPort defines the port of the gui server, when using
// `govuegui.Serve()`
var ServerPort = ":2700"

// NewRouter returns a router from the gorillatoolkit
// http://www.gorillatoolkit.org/pkg/mux
// The router already includes all the paths which are needed
// for the gui. It can be called like:
//   r := govuegui.NewRouter()
//   r.HandleFunc("/products/{key}", ProductHandler)
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix+"/vue.min.js", vuejsdev.Handler)
	return r
}

// Serve wraps the http.ListenAndServe() function, but adds the
// routes for the gui.
func Serve() error {
	r := NewRouter()
	return http.ListenAndServe(ServerPort, r)
}