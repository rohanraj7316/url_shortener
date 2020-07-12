package routers

import (
	"url_shortener/app/handlers"

	"github.com/go-chi/chi"
)

// CreateRouter create short urls
func CreateRouter(r chi.Router) {
	r.Get("/health", h(handlers.GetHealth))
}
