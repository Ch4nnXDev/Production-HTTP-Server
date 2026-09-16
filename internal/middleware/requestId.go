package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		requestID := uuid.New().String()

		ctx := context.WithValue(
			req.Context(),
			requestIDKey,
			requestID,
		)

		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}