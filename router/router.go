package router

import (
	"net/http"
	"urlshort/api"
	)

type Router struct {
	routes map[string]map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.Handler),
	}
}

func (r *Router) registerRoute(path, method string, handler http.Handler) {
	if _, ok := r.routes[path]; !ok {
		r.routes[path] = make(map[string]http.Handler)
	}
	r.routes[path][method] = handler
}

func (r *Router) GET(path string, handler http.Handler) {
	r.registerRoute(path, http.MethodGet, handler)
}

func (r *Router) POST(path string, handler http.Handler) {
	r.registerRoute(path, http.MethodPost, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    pathHandlers, pathFound := r.routes[req.URL.Path]
    if !pathFound {
		api.RespondWithError(w, http.StatusNotFound, "Not Found")
        return
    }

    handler, methodFound := pathHandlers[req.Method]
    if !methodFound {
		api.RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
        return
    }

    handler.ServeHTTP(w, req)
}