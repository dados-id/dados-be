package api

import (
	"firebase.google.com/go/auth"
	"github.com/dados-id/dados-be/config"
	db "github.com/dados-id/dados-be/db/sqlc"
	"github.com/dados-id/dados-be/exception"
	"github.com/dados-id/dados-be/util"
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

	if server.config.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.New()
	router.Use(util.HttpLogger())
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
		userRoutes.GET("/users", server.getUser)
		userRoutes.PUT("/users", server.updateUser)
		userRoutes.GET("/users/professor_ratings", server.userListProfessorRatings)
		userRoutes.GET("/users/school_ratings", server.userListSchoolRatings)
		userRoutes.GET("/users/saved_professors", server.userListSavedProfessors)
		userRoutes.DELETE("/users/professors/:professor_id", server.unsaveProfessor)
		userRoutes.POST("/users/professors/:professor_id", server.saveProfessor)

		schoolRoutes := authRoutes.Group("/")
		schoolRoutes.POST("/schools", server.createSchool)
		schoolRoutes.GET("/schools/:school_id", server.getSchoolInfo)
		schoolRoutes.GET("/schools", server.listSchools)
		schoolRoutes.PUT("/schools/:school_id", server.updateSchoolStatusRequest)

		schoolRatingRoutes := authRoutes.Group("/")
		schoolRatingRoutes.GET("/schools/:school_id/ratings/:school_rating_id", server.getSchoolRating)
		schoolRatingRoutes.GET("schools/:school_id/ratings", server.listSchoolRatings)
		schoolRatingRoutes.POST("schools/:school_id/ratings", server.createSchoolRating)
		schoolRatingRoutes.PUT("schools/:school_id/ratings/:school_rating_id", server.updateSchoolRating)

		professorRoutes := authRoutes.Group("/")
		professorRoutes.POST("/professors", server.createProfessor)
		professorRoutes.GET("/professors/:professor_id", server.getProfessorInfo)
		professorRoutes.GET("/professors", server.listProfessors)
		professorRoutes.GET("schools/:school_id/professors", server.listProfessorsBySchool)
		professorRoutes.GET("schools/:school_id/faculties/:faculty_id/professors", server.listProfessorsBySchoolAndFaculty)
		professorRoutes.PUT("/professors/:professor_id", server.updateProfessorStatusRequest)

		professorRatingRoutes := authRoutes.Group("/")
		professorRatingRoutes.GET("/professors/:professor_id/ratings/:professor_rating_id", server.getProfessorRating)
		professorRatingRoutes.GET("professors/:professor_id/ratings", server.listProfessorRatings)
		professorRatingRoutes.POST("professors/:professor_id/ratings", server.createProfessorRating)
		professorRatingRoutes.PUT("professors/:professor_id/ratings/:professor_rating_id", server.updateProfessorRating)

		facultyRoutes := authRoutes.Group("/")
		facultyRoutes.GET("schools/:school_id/faculties", server.listFacultiesBySchool)
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
