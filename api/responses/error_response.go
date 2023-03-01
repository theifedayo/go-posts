package responses

import (
	"encoding/json"
	"net/http"
)

// represents the response sent in case of errors.
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: errorMessage})
}