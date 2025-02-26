package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router is a struct that represents a route of the API
type Router struct {
	URI             string
	Method          string
	Function        func(w http.ResponseWriter, r *http.Request)
	IsAutentication bool
}

func RouterConfig(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes{
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}