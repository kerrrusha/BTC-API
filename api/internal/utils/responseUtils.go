package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	CheckForError(err)
}
