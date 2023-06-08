package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type Server struct {
	ginEngine  *gin.Engine
	httpServer *http.Server
}

func NewServer() *Server {
	route := gin.Default()
	return &Server{
		ginEngine: route,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: route,
		},
	}
}

func (s *Server) SetRoutes(routes []Route) {
	for _, route := range routes {
		s.ginEngine.Handle(route.Method, route.Path, route.Handler)
	}

}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) GinEngine() *gin.Engine {
	return s.ginEngine
}
