package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"github.com/lohyangxian/OneCV-Govtech/internal/api"
	"github.com/lohyangxian/OneCV-Govtech/internal/database"
	"github.com/lohyangxian/OneCV-Govtech/internal/services"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"os"
)

func main() {

	// Load the configuration
	configuration := loadConfig()

	// Overwrite configuration values with environment variables
	overwriteConfig(&configuration)

	// Open a connection to the database
	db := database.ConnectToDB(configuration)

	// Initialize services
	studentService := services.NewStudentService(db)
	teacherService := services.NewTeacherService(db)

	// Initialize and start the server
	startServer(studentService, teacherService, &configuration, db)
}

func loadConfig() config.Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var configuration config.Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode config into struct, %v", err)
	}

	viper.SetDefault("database.dbname", "govtech_db_test")

	return configuration
}

func overwriteConfig(configuration *config.Configurations) {
	// Check if environment variables are set and update configuration values accordingly
	if envDBHost := os.Getenv("DB_HOST"); envDBHost != "" {
		configuration.Database.DBHost = envDBHost
	}
	if envDBUser := os.Getenv("DB_USER"); envDBUser != "" {
		configuration.Database.DBUser = envDBUser
	}
	if envDBPassword := os.Getenv("DB_PASSWORD"); envDBPassword != "" {
		configuration.Database.DBPassword = envDBPassword
	}
	if envDBName := os.Getenv("DB_NAME"); envDBName != "" {
		configuration.Database.DBName = envDBName
	}
	if envDBPort := os.Getenv("DB_PORT"); envDBPort != "" {
		configuration.Database.Port = envDBPort
	}
}

func startServer(studentService services.StudentService, teacherService services.TeacherService, configuration *config.Configurations, db *gorm.DB) {
	server, err := api.NewServer(studentService, teacherService, configuration, db)
	if err != nil {
		fmt.Printf("Error starting server, %v", err)
	}

	port := configuration.Server.Port
	var address string
	if port == "" {
		address = ":3000"
	} else {
		address = ":" + port
	}

	err = server.Start(address)
	if err != nil {
		fmt.Printf("Error starting server, %v", err)
	}
}
