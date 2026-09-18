package middleware

import (
	"net/http"
	"log"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	
}
//* = means this points to the type of this variable
//& = means this points to the variable address

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(rw, req)

		requestID := req.Context().Value(requestIDKey)
		

		log.Println(requestID, req.Method, req.URL.Path, rw.statusCode, time.Since(start), req.RemoteAddr)

		

	})

}

func (rw *responseWriter) WriteHeader(statusCode int) {

	if rw.statusCode != 0 {
		return
	}
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}


func (rw *responseWriter) Write(body []byte) (int, error) {

    if rw.statusCode == 0 {
        rw.WriteHeader(http.StatusOK)
    }

    return rw.ResponseWriter.Write(body)
}