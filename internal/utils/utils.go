package utils

import (
	"encoding/json"
	"net/http"
)

// sending a JSON response with a given status code and payload
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
