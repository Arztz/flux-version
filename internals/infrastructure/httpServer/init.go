package httpServer

import (
	"flux-version/internals/config"
	"net/http"
)

type Server struct {
	Config  config.Configuration
	HttpMux *http.ServeMux
}

func NewServer(config config.Configuration, httpMux *http.ServeMux) *Server {
	s := &Server{
		Config:  config,
		HttpMux: httpMux,
	}
	//s.Configure(context.Background(), opts)
	return s
}
