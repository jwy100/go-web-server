package server

import (
	"context"
	"fmt"
	"net/http"
)

type WebServer interface {
	Serve()
	Stop()
	Config() WebServerConfig
}

type StaticWebServer struct {
	config WebServerConfig
	server http.Server
}

func NewStaticWebServer(config WebServerConfig) StaticWebServer {
	return StaticWebServer{
		config: config,
	}
}

func (s StaticWebServer) Serve() {
	mux := http.NewServeMux()

	mux.HandleFunc(s.config.ContextPath(), s.BaseHandler)

	srv := &http.Server{
		Addr:    ":" + s.config.Port(),
		Handler: mux,
	}
	srv.ListenAndServe()
}

func (s StaticWebServer) BaseHandler(w http.ResponseWriter, r *http.Request) {
	// use a strategy to get the handler func given a request,
	// then use the handler func and return a response
	path := r.URL.Path
	method := r.Method
	for _, route := range s.config.Routes() {
		if route.Path() == path && contains(route.HttpMethods(), method) {
			route.handleFunc(w, r)
			return
		}
	}
	fmt.Fprintf(w, "404 Not Found")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s StaticWebServer) Config() WebServerConfig {
	return s.config
}

func (s StaticWebServer) Stop() {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
