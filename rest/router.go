package rest

const prefix = "/v1"

// setupRoutes sets up the routes for the server
func (s *Server) setupRoutes() {
	v1 := s.Engine.Group(prefix)
	v1.GET("/health", s.healthHandler)
}
