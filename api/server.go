package api

import (
	"firebase.google.com/go/auth"
	"github.com/dados-id/dados-be/config"
	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests.
type Server struct {
	config         config.Config
	query          db.Querier
	firebaseClient auth.Client
	router         *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(configuration config.Config, query db.Querier, firebaseClient auth.Client) (*Server, error) {
	server := &Server{
		config:         configuration,
		query:          query,
		firebaseClient: firebaseClient,
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

	router.POST("/users/login", server.loginUser)
	router.POST("/users", server.createUser)

	authRoutes := router.Group("/")
	authRoutes.Use()
	{
		userRoutes := authRoutes.Group("/")
		userRoutes.GET("/users/:id", server.getUser)
		userRoutes.PUT("/users/:id", server.updateUser)

		userRoutes.GET("/users/:id/professor_ratings", server.userListProfessorRatings)
		userRoutes.GET("/users/:id/school_ratings", server.userListSchoolRatings)
		userRoutes.GET("/users/:id/saved_professors", server.userListSavedProfessors)

		userRoutes.DELETE("/users/:user_id/professors/:professor_id", server.unsaveProfessor)
		userRoutes.POST("/users/:user_id/professors/:professor_id", server.saveProfessor)

		schoolRoutes := authRoutes.Group("/")
		schoolRoutes.POST("/schools", server.createSchool)
		schoolRoutes.GET("/schools/:school_id", server.getSchoolInfoAggregate)
		schoolRoutes.GET("/schools", server.listSchools)
		schoolRoutes.PUT("/schools/:school_id", server.updateSchoolStatusRequest)

		schoolRatingRoutes := authRoutes.Group("/")
		schoolRatingRoutes.GET("/schools/:school_id/ratings/:school_rating_id", server.getSchoolRating)
		schoolRatingRoutes.GET("schools/:school_id/ratings", server.listSchoolRatings)
		schoolRatingRoutes.POST("schools/:school_id/ratings", server.createSchoolRating)
		schoolRatingRoutes.PUT("schools/:school_id/ratings/:school_rating_id", server.updateSchoolRating)
	}

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) start(address string) error {
	return server.router.Run(address)
}

func RunGinServer(configuration config.Config, query db.Querier, firebaseClient auth.Client) {
	server, err := NewServer(configuration, query, firebaseClient)
	exception.FatalIfNeeded(err, "cannot create server")

	err = server.start(configuration.HTTPServerAddress)
	exception.FatalIfNeeded(err, "cannot start server")
}
