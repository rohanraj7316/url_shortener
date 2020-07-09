package handlers

import (
	"net/http"
)

// Handler generic structure for all handlers
type Handler func(w http.ResponseWriter, r *http.Request) error

// ServerRequestHandler generic error handler
func (fn Handler) ServerRequestHandler(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		// TODO: need to add logging
	}
}
