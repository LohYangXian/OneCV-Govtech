package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"github.com/lohyangxian/OneCV-Govtech/internal/services"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	StudentService services.StudentService
	TeacherService services.TeacherService
	Config         *config.Configurations
	Database       *gorm.DB
	Router         *gin.Engine
}

func NewServer(studentService services.StudentService, teacherService services.TeacherService, config *config.Configurations, database *gorm.DB) (*Server, error) {
	server := &Server{
		StudentService: studentService,
		TeacherService: teacherService,
		Config:         config,
		Database:       database,
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

	port := s.Config.Server.Port
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

	s.Router = router
}

func (s *Server) Start(address string) error {
	return s.Router.Run(address)
}
