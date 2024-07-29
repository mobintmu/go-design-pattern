package main

import (
	"fmt"
	"net/http"
	"time"
)

const port = ":4000"

type application struct{}

func main() {
	app := application{}
	server := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting server on port", port)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
