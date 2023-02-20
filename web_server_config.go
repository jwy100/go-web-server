package server

type WebServerConfig struct {
	port        string
	contextPath string
	routes      []Route
}

func NewWebServerConfig(port string, contextPath string, routes []Route) WebServerConfig {
	return WebServerConfig{
		port:        port,
		contextPath: contextPath,
		routes:      routes,
	}
}

func (w WebServerConfig) Port() string {
	return w.port
}

func (w WebServerConfig) ContextPath() string {
	return w.contextPath
}

func (w WebServerConfig) Routes() []Route {
	return w.routes
}
