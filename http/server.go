package http

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	return &Server{router: router}
}

func (s *Server) Start() {
	s.RegisterNodesRoutes()
	s.RegisterTasksRoutes()
	s.RegisterTailscaleRoutes()

	s.router.Run()
}