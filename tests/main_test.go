package tests

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"github.com/lohyangxian/OneCV-Govtech/internal/api"
	"github.com/lohyangxian/OneCV-Govtech/internal/services"
	"github.com/lohyangxian/OneCV-Govtech/tests/mocks"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
)

//go:embed test_db/000001_create_empty_tables.up.sql
var createEmptyTablesSQL string

//go:embed test_db/000001_create_empty_tables.down.sql
var dropTablesSQL string

//go:embed test_db/000002_seed_students_and_teachers.sql
var seedStudentsTeachersSQL string

var testDB *gorm.DB
var err error
var studentService services.StudentService
var teacherService services.TeacherService
var configurations config.Configurations

func TestMain(m *testing.M) {

	// Load the configuration
	configuration := loadTestConfig()

	// Overwrite configuration values with environment variables
	overwriteConfig(&configuration)

	// Set up the test database
	testDB, err = SetUpTestDBConnection(configuration)
	if err != nil {
		fmt.Println("failed to connect to the test database", err)
		os.Exit(1)
	}

	// Execute SQL script to seed test data
	SetUpTestDB(testDB)
	studentService = services.NewStudentService(testDB)
	teacherService = services.NewTeacherService(testDB)

	// Run the tests
	code := m.Run()

	// Tear down the test database
	TearDownTestDB(testDB)

	// Exit with the result of the tests
	os.Exit(code)
}

func SetUpMockServer(testDB *gorm.DB) (*gin.Engine, *api.Server) {
	router := gin.New()

	mockTeacherService := &mocks.MockTeacherService{}
	mockStudentService := &mocks.MockStudentService{}

	mockServer := &api.Server{
		StudentService: mockStudentService,
		TeacherService: mockTeacherService,
		Config:         &configurations,
		Database:       testDB,
		Router:         router,
	}

	return router, mockServer
}

func SetUpTestDBConnection(configuration config.Configurations) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configuration.TestDatabase.TestDBHost,
		configuration.TestDatabase.TestDBUser,
		configuration.TestDatabase.TestDBPassword,
		configuration.TestDatabase.TestDBName,
		configuration.TestDatabase.TestPort,
	)

	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to the test database", err)
		return db, err
	}

	return db, nil
}

func SetUpTestDB(db *gorm.DB) {
	//Execute SQL script to create test data
	db = db.Exec(seedStudentsTeachersSQL)
}

func TearDownTestDB(db *gorm.DB) {
	//Execute SQL script to delete test data
	db = db.Exec(dropTablesSQL)
}

func loadTestConfig() config.Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
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
	if envDBHost := os.Getenv("TEST_DB_HOST"); envDBHost != "" {
		configuration.TestDatabase.TestDBHost = envDBHost
	}
	if envDBUser := os.Getenv("TEST_DB_USER"); envDBUser != "" {
		configuration.TestDatabase.TestDBUser = envDBUser
	}
	if envDBPassword := os.Getenv("TEST_DB_PASSWORD"); envDBPassword != "" {
		configuration.TestDatabase.TestDBPassword = envDBPassword
	}
	if envDBName := os.Getenv("TEST_DB_NAME"); envDBName != "" {
		configuration.TestDatabase.TestDBName = envDBName
	}
	if envDBPort := os.Getenv("TEST_DB_PORT"); envDBPort != "" {
		configuration.TestDatabase.TestPort = envDBPort
	}
}
