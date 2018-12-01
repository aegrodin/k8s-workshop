package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting app...")

	port := "8080"

	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is stoped with error: %v", err)
	}

	log.Print("App has been stopped")
}
