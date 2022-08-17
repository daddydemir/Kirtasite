package main

import (
	"Kirtasite/config"
	"Kirtasite/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:1337/api/v1/")

	config.GetConnect()

	server := &http.Server{
		Addr:    ":1337",
		Handler: handlers.MainRouting(),
	}

	server.ListenAndServe()
}
