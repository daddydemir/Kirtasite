package main

import (
	"Kirtasite/config"
	"Kirtasite/handlers"
	"log"
	"net/http"
)

func main() {
	println("http://localhost" + config.Get("SERVER_PORT") + "/api/v1/")

	config.GetConnect()
	config.RedisClient()

	server := &http.Server{
		Addr:    config.Get("SERVER_PORT"),
		Handler: handlers.MainRouting(),
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
