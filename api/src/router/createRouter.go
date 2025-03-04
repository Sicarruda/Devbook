package router

import "github.com/gorilla/mux"

// CreateRouter is a function that returns a new mux.Router
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return RouterConfig(r)
}
