package main

import (
	"url_shortener/app/services"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			// TODO: need to think abt the ways to handle this.
		}
	}()

	// TODO: ways to handle .env files
	// TODO: call to services.

}

func h(fn services.Service) bool {
	return services.Service(fn).StartServer
}
