package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/chetan177/oneclick/rest/models"
	"github.com/chetan177/oneclick/version"
)

// healthHandler handles the /health route
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthResponse{
		Status: "healthy",
		Version: version.GetVersion(),
	})
}
