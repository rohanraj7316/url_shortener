package services

import (
	"net/http"
	"url_shortener/app/routers"
)

// StartServer initializing and starting http server
func StartServer() error {

	mux := routers.RouterHandler()
	server := http.Server{
		Addr:    ":2401",
		Handler: mux,
	}

	panic(server.ListenAndServe())
}
