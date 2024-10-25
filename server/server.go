package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/go-crud/middlewares"
	"github.com/yosefalemu/go-crud/routes"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (s *Server) Initialize() {
	s.Router.Use(middlewares.CORSMiddleware())
	routes.RegisterRoutes(s.Router)
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
