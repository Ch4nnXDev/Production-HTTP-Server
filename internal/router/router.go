package router

import "net/http"

type Router struct {
	routes map[string]http.Handler
}

func NewRouter() *Router {

	return &Router {

		routes: make(map[string]http.Handler),
	}
}

func (r *Router) Handle(path string, handler http.Handler, method string) {
	route := method + " " + path
	r.routes[route] = handler
	
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := req.Method + " " + req.URL.Path
	handler, ok := r.routes[route]

	if !ok {
		http.NotFound(w, req)
		return
	}

	handler.ServeHTTP(w, req)
}