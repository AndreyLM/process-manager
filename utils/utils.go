package utils

import (
	"encoding/json"
	"net/http"
)

// PrepareData - returns data for respose
func PrepareData(code int, message string, response interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Code":     code,
		"Message":  message,
		"Response": response,
	}
}

// Respond - responses with given data
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
