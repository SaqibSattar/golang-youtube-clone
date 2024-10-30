package common

import (
	"encoding/json"
	"net/http"
)

// ApiResponse represents a standardized API response
type ApiResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	Success    bool        `json:"success"`
}

// ApiError represents a standardized error response
type ApiError struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors,omitempty"`
	Success    bool     `json:"success"`
}

// Respond sends a JSON response with the given status code and data/message
func Respond(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ApiResponse{
		StatusCode: statusCode,
		Data:       data,
		Message:    message,
		Success:    statusCode < 400,
	}
	json.NewEncoder(w).Encode(response)
}

// RespondError sends a standardized error response
func RespondError(w http.ResponseWriter, err ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}
