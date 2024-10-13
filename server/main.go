package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hi banana!")
	server := http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadHeaderTimeout: time.Duration(3) * time.Second,
	}
	log.Printf("Serving on port: %s\n", ":8080")
	log.Fatal(server.ListenAndServe())
}
