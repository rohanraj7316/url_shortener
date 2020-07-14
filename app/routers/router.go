package routers

import (
	"net/http"
	"url_shortener/app/handlers"

	"github.com/go-chi/chi"
)

// RouterHandler generic router handler
func RouterHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", RedirectRouter)
	r.Route("/url-shortener", CreateRouter)
	return r
}

// check for all handler
func h(fn handlers.Handler) http.HandlerFunc {
	return handlers.Handler(fn).ErrorHandler
}
