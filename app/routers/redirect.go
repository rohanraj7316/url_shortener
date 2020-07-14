package routers

import (
	"url_shortener/app/handlers/redirect"

	"github.com/go-chi/chi"
)

// RedirectRouter url redirection
func RedirectRouter(r chi.Router) {
	r.Get("/{urlId}", h(redirect.Redirect))
}
