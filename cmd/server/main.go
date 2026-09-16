package main

import (
	"log"
	"net/http"

	"github.com/channakarawita/production-http-server/internal/middleware"
	"github.com/channakarawita/production-http-server/internal/router"
)

func main() {

	r := router.NewRouter()

	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Method:", req.Method)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GO HTTP SERVER"))
	}), "GET")

	r.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	}), "GET")

	handler := middleware.RequestID(
		middleware.Logging(r),
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server Listening On: 8080")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}