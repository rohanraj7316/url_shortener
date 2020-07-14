package routers

import (
	"url_shortener/app/handlers"
	"url_shortener/app/handlers/create"

	"github.com/go-chi/chi"
)

// CreateRouter create short urls
func CreateRouter(r chi.Router) {
	r.Get("/health", h(handlers.GetHealth))
	r.Post("/create", h(create.ShortURL))
	// r.Patch("/update", h())
}
