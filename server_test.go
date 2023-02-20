package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func healthEndpointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
func TestHealthEndpointReturnsOk(t *testing.T) {

	s := NewStaticWebServer(WebServerConfig{port: "8090", contextPath: "/", routes: []Route{NewRoute("/health", []string{"GET"}, healthEndpointHandler)}})
	go s.Serve()
	defer s.Stop()

	res, err := http.Get("http://localhost:8090/health")
	if err != nil {
		t.Fatal("HTTP request failed")
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	if string(resBody) != "OK" {
		t.Fatalf("Expected OK, but actual was %s", resBody)
	}
}
