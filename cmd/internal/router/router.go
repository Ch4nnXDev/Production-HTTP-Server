package router

import "net/http"

type Router struct {
	routes map[string]http.Handler
}

func NewRouter() *Router {
	router := Router{}

	return &router
}