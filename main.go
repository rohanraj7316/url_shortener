package main

import (
	"log"
	"url_shortener/app/services"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			// TODO: need to think abt the ways to handle this.
		}
	}()

	log.Println("starting the server")
	// load and intiate all the services
	services.InitService()
}
