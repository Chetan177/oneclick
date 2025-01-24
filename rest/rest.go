package rest

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/chetan177/oneclick/db"
)

const (
	envPort = "PORT"
)

// Server holds the Gin engine and Kubernetes client
type Server struct {
	Echo  *echo.Echo
	Port  string
	DB    *db.DB
	Log   echo.Logger
	Token string
	Validate *validator.Validate
}


// NewServer initializes and returns a new Server instance
func NewServer() *Server {
	port := os.Getenv(envPort)
	if port == "" {
		port = ":8000"
	} else {
		port = ":" + port
	}
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(RequestLoggerMiddleware)

	db, err := db.NewDB(os.Getenv("MONGO_URI"))
	if err != nil {
		server.Logger.Fatal(err)
	}

	return &Server{
		Echo: server,
		Port: port,
		DB:   db,
		Log:  server.Logger,
		Validate: validator.New(),
	}
}


func RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Process request
		err := next(c)

		// Log request and response details
		req := c.Request()
		res := c.Response()
		responseTime := time.Since(start)

		c.Logger().Infof("Request: %s %s, Response: %d, Time: %v", req.Method, req.URL, res.Status, responseTime)

		return err
	}
}

func (s *Server) TokenAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiToken := c.Request().Header.Get("API_TOKEN")
			if apiToken != s.Token {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API token")
			}
			return next(c)
		}
	}
}

func GenerateToken() string {
	length := 24
	bytes := make([]byte, length)

	// Fill the byte slice with random data
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	// Convert the byte slice to a hex string
	return hex.EncodeToString(bytes)
}

func (s *Server) Start() {
	s.setupRoutes()
	go func() {
		s.Echo.Start(s.Port)
	}()

	// err := s.DB.CreateOrUpdateSystemData(token)
	token, err := s.DB.CreateOrGetToken(GenerateToken())
	if err != nil {
		s.Log.Fatal(err)
	}
	// Update the token in the server
	s.Token = token
	fmt.Printf("****************************************************************** \n \n")
	fmt.Printf("🚀 API Token: %s \n \n", token)
	fmt.Printf("****************************************************************** \n \n")

}

func (s *Server) Stop() {
	s.Log.Info("Stopping server")
	s.Echo.Shutdown(context.Background())
}

func (s *Server) ValidateRequest(c echo.Context, request interface{}) error {
	return s.Validate.Struct(request)
}
