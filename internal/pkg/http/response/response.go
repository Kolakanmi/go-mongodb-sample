package response

import (
	"encoding/json"
	"net/http"
)

type (
	responseBody struct {
		Success bool `json:"success"`
		Message string `json:"message,omitempty"`
		Data interface{} `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}
	
	response struct {
		Body *responseBody
		StatusCode int
	}
)

func (r response) ToJSON(w http.ResponseWriter) error{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(r.StatusCode)
	return json.NewEncoder(w).Encode(r.Body)
}

func newResponse(success bool, message string, err error, data interface{}, status int) *response {
	return &response{
		Body:       &responseBody{
			Success: success,
			Message: message,
			Data:    data,
			Error: err.Error(),
		},
		StatusCode: status,
	}
}

func OK(message string, data interface{}) *response {
	return newResponse(true, message, nil, data, http.StatusOK)
}

func Fail(err error, status int) *response {
	return newResponse(false,"", err, nil, status)
}
