package router

import "net/http"

type Router struct {
	routes map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.Handler),
	}
}

func (r *Router) Handle(path string, handler http.Handler) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if hanlder, ok := r.routes[path]; ok {
		hanlder.ServeHTTP(w, req)
		return
	}

	http.NotFound(w, req)
}