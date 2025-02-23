package router

import "net/http"

// Router is a struct that represents a route of the API
type Router struct {
	URI             string
	Method          string
	Function        func(w http.ResponseWriter, r *http.Request)
	IsAutentication bool
}
