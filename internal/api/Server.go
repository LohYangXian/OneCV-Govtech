package api

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"log"
)

// TODO: Check if gorm.DB is better than sql.DB
type Server struct {
	config   *config.Configurations
	database *sql.DB
	router   *gin.Engine
}

func NewServer(config *config.Configurations, database *sql.DB) (*Server, error) {
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

	port := s.config.Server.Port
	var address string
	if port == "" {
		address = ":3000"
	} else {
		address = ":" + port
	}

	err := router.Run(address)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
