package server

import (
	"context"
	"net/http"
)

type WebServer interface {
	Serve() error
	Stop()
}

type StaticWebServer struct {
	server http.Server
}

func NewStaticWebServer(port string, router Router) StaticWebServer {

	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return StaticWebServer{
		server: srv,
	}
}

func (s StaticWebServer) Serve() error {
	return s.server.ListenAndServe()
}

func (s StaticWebServer) Stop() {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}
