package main

import (
	"log"
	"net/http"
)

func main() {

	ap, err := initializeDI()
	if err != nil {
		panic(err)
	}

	s := http.Server{
		Addr:    ":8080",
		Handler: ap.route,
	}

	log.Fatal(s.ListenAndServe())
}
