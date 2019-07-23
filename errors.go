package gokitkit

import (
	"context"
	"encoding/json"
	"net/http"
)

// ErrorWrapper
type ErrorWrapper struct {
	Error string `json:"error"`
}

// ErrorEncoder
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
