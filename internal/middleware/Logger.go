package middleware

import (
	"net/http"
	"log"
)



func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		log.Println(req.Method, req.URL.Path)

		next.ServeHTTP(w, req)

	})

}