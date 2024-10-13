package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi banana!")
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	log.Printf("Serving on port: %s\n", ":8080")
	log.Fatal(server.ListenAndServe())
}
