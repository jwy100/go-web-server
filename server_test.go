package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestHealthEndpointReturnsOk(t *testing.T) {

	s := NewStaticWebServer(WebServerConfig{port: "8090", contextPath: "/api", routes: GetRoutes()})
	go s.Serve()
	defer s.Stop()

	res, err := http.Get("http://localhost:8090/api/health")
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

func TestMetricsEndpointReturnsMetrics(t *testing.T) {

	s := NewStaticWebServer(WebServerConfig{port: "8090", contextPath: "/api", routes: GetRoutes()})
	go s.Serve()
	defer s.Stop()

	res, err := http.Get("http://localhost:8090/api/metrics")
	if err != nil {
		t.Fatal("HTTP request failed")
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	if string(resBody) != "Metrics" {
		t.Fatalf("Expected Metrics, but actual was %s", resBody)
	}
}
