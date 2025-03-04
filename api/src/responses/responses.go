package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Json returns a json response to the request
func Json(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error returns an error response to the request in json format
func Error(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})

}
