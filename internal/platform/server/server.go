package server

import (
	"github.com/Frankcs96/learning-platform-service/internal/platform/handler/health"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port string) Server {
	srv := Server{
		httpAddr: host + ":" + port,
		engine:   gin.New(),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
