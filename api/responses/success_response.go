package responses

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse represents the response sent in case of successful requests.
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// NewSuccessResponse creates a new SuccessResponse instance with the specified data.
func NewSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(SuccessResponse{Data: data})
}