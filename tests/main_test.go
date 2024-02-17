package tests

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"github.com/lohyangxian/OneCV-Govtech/internal/api"
	"github.com/lohyangxian/OneCV-Govtech/tests/mocks"
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

func TestMain(m *testing.M) {
	// Set up the test database
	testDB, err = SetUpTestDBConnection()
	if err != nil {
		fmt.Println("failed to connect to the test database", err)
		os.Exit(1)
	}

	// Execute SQL script to seed test data
	SetUpTestDB(testDB)

	// Run the tests
	code := m.Run()

	// Tear down the test database
	TearDownTestDB(testDB)

	// Exit with the result of the tests
	os.Exit(code)
}

func SetUpMockServer(testDB *gorm.DB) (*gin.Engine, *api.Server) {
	router := gin.New()

	mockConfig := &config.Configurations{
		Server: config.ServerConfig{
			Port: "3000",
		},
		Database: config.DatabaseConfig{
			DBUser:     "root",
			DBPassword: "root123",
			DBName:     "govtechdb_test",
			Port:       "5432",
		},
		Environment: "test",
	}

	mockTeacherService := &mocks.MockTeacherService{}
	mockStudentService := &mocks.MockStudentService{}

	mockServer := &api.Server{
		StudentService: mockStudentService,
		TeacherService: mockTeacherService,
		Config:         mockConfig,
		Database:       testDB,
		Router:         router,
	}

	return router, mockServer
}

func SetUpTestDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=root password=root123 dbname=govtechdb_test port=5432 sslmode=disable"), &gorm.Config{})
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
