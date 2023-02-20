package server

import "net/http"

type Route struct {
	path        string
	httpMethods []string
	handleFunc  func(w http.ResponseWriter, r *http.Request)
}

func NewRoute(path string, httpMethods []string, handleFunc http.HandlerFunc) Route {
	return Route{
		path:        path,
		httpMethods: httpMethods,
		handleFunc:  handleFunc,
	}
}

func (r Route) Path() string {
	return r.path
}

func (r Route) HttpMethods() []string {
	return r.httpMethods
}

func (r Route) HandleFunc() func(w http.ResponseWriter, r *http.Request) {
	return r.handleFunc
}
