package routers

import (
	"url_shortener/app/handlers"

	"github.com/go-chi/chi"
)

// RedirectRouter url redirection
func RedirectRouter(r chi.Router) {
	r.Post("/health", h(handlers.GetHealth))
}
