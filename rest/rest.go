package rest

import (
	"github.com/gin-gonic/gin"
)

// Server holds the Gin engine
type Server struct {
	Engine *gin.Engine
}

// NewServer initializes and returns a new Server instance
func NewServer() *Server {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	server := &Server{Engine: engine}

	// Set up routes
	server.setupRoutes()

	return server
}

func (s *Server) Start() {
	s.Engine.Run(":8080")
}
