package main

import (
	"log"
	"net/http"
	"os"
	"context"
	"time"
	"os/signal"
	"syscall"

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

	r.Handle("/panic", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		panic("this is a test panic")
	}), "GET")

	handler := middleware.RequestID(
		middleware.Recovery(
			middleware.Logging(r),
		),
		
	)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Println("Server Listening On: 8080")

	go shutdown(server)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}


func shutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) //first param is the channel name 
	defer signal.Stop(quit)
	<-quit
	log.Println("shutdown signal received")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("Server shutdown error:", err)
	}

	log.Println("Server shutdown complete")
}