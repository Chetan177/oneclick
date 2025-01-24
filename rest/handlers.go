package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/chetan177/oneclick/version"
	"github.com/chetan177/oneclick/models"
)

// healthHandler handles the /health route
func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthResponse{
		Status: "healthy",
		Version: version.GetVersion(),
	})
}

func (s *Server) tokenHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthResponse{
		Status: "Token Verified",
		Version: s.Token,
	})
}

