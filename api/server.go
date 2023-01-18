package api

import (
	"github.com/dados-id/dados-be/config"
	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests.
type Server struct {
	config config.Config
	query  db.Querier
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(configuration config.Config, query db.Querier) (*Server, error) {
	server := &Server{
		config: configuration,
		query:  query,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	if server.config.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(HttpLogger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/users", server.createUser)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) start(address string) error {
	return server.router.Run(address)
}

func RunGinServer(configuration config.Config, query db.Querier) {
	server, err := NewServer(configuration, query)
	exception.FatalIfNeeded(err, "cannot create server")

	err = server.start(configuration.HTTPServerAddress)
	exception.FatalIfNeeded(err, "cannot start server")
}
