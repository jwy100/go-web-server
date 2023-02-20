package server

import "net/http"

type Router struct {
	routes []Route
}

func (r Router) Routes() []Route {
	return r.routes
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	for _, route := range r.routes {
		if route.Path() == path && route.HttpMethod() == method {
			route.handleFunc(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
