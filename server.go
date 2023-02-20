package server

import (
	"context"
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
	path := r.URL.Path
	method := r.Method
	for _, route := range s.config.Routes() {
		if route.Path() == path && route.HttpMethod() == method {
			route.handleFunc(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func (s StaticWebServer) Config() WebServerConfig {
	return s.config
}

func (s StaticWebServer) Stop() {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}
