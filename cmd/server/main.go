package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method:", r.Method)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GO HTTP SERVER"))
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server Listening On: 8080")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}