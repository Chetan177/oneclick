package rest

import (
	"github.com/labstack/echo/v4"
)

const prefix = "/v1"

// setupRoutes sets up the routes for the server
func (s *Server) setupRoutes() {
	s.Echo.GET("/health", s.healthHandler)
	v1 := s.Echo.Group(prefix)
	v1.Use(s.TokenAuthMiddleware())
	v1.GET("/auth/token", s.tokenHandler)
	s.setupProjectRoutes(v1)
}

func (s *Server) setupProjectRoutes(v1 *echo.Group) {
	v1.POST("/project", s.createProjectHandler)
	v1.GET("/project", s.getProjectsHandler)
	v1.GET("/project/:id", s.getProjectHandler)
	v1.PUT("/project/:id", s.updateProjectHandler)
	v1.DELETE("/project/:id", s.deleteProjectHandler)
}

