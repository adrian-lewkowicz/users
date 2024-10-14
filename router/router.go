package router

import (
	"net/http"
	"strings"
)

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{make(map[string]http.HandlerFunc)}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if handler, ok := r.routes[path]; ok {
		handler(w, req)
		return
	}

	for route, handler := range r.routes {
		if strings.Contains(route, "{id}") {
			baseRoute := strings.Split(route, "/{id}")[0]
			if strings.HasPrefix(path, baseRoute) {
				handler(w, req)
				return
			}
		}
	}

	http.NotFound(w, req)
}

func (r *Router) Handle(route string, handler http.HandlerFunc) {
	r.routes[route] = handler
}
