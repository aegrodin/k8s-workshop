package main

import (
	"fmt"
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

	http.HandleFunc("/", handler)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is stoped with error: %v", err)
	}

	log.Print("App has been stopped")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
