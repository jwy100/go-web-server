package server

import "net/http"

type Route struct {
	path       string
	httpMethod string
	handleFunc func(w http.ResponseWriter, r *http.Request)
}

func NewRoute(path string, httpMethod string, handleFunc http.HandlerFunc) Route {
	return Route{
		path:       path,
		httpMethod: httpMethod,
		handleFunc: handleFunc,
	}
}

func (r Route) Path() string {
	return r.path
}

func (r Route) HttpMethod() string {
	return r.httpMethod
}

func (r Route) HandleFunc() func(w http.ResponseWriter, r *http.Request) {
	return r.handleFunc
}
