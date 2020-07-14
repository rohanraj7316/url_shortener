package main

import (
	"net/http"
	"url_shortener/app/routers"
	"url_shortener/app/services"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			// TODO: need to think abt the ways to handle this.
		}
	}()

	// load and intiate all the services
	services.InitService()

	// load routers
	mux := routers.RouterHandler()

	// load http servers
	server := http.Server{
		Addr:    ":2401",
		Handler: mux,
	}
	panic(server.ListenAndServe())
}
