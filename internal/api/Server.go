package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	config   *Config
	database *Database
}

func NewServer(config *Config, database *Database) (*Server, error) {
	server := &Server{
		config:   config,
		database: database,
	}

	server.initRoutes()
	return server, nil
}

func (s *Server) initRoutes() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	api.POST("/register", s.Register)
	api.GET("/commonstudents", s.CommonStudents)
	api.POST("/suspend", s.Suspend)
	api.POST("/retrievefornotifications", s.RetrieveForNotifications)

	err := router.Run(s.config.Address)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
