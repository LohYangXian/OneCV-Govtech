package database

import (
	_ "embed"
	"fmt"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/000001_create_empty_tables.up.sql
var createEmptyTablesSQL string

//go:embed migrations/000001_create_empty_tables.down.sql
var dropTablesSQL string

//go:embed migrations/000002_seed_students_and_teachers.sql
var seedStudentsTeachersSQL string

func ConnectToDB(configuration config.Configurations) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configuration.Database.DBHost,
		configuration.Database.DBUser,
		configuration.Database.DBPassword,
		configuration.Database.DBName,
		configuration.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database, %v", err)
		return nil
	}

	CreateEmptyTables(db)

	return db
}

func CreateEmptyTables(db *gorm.DB) {
	//Execute SQL script to create empty tables
	db = db.Exec(createEmptyTablesSQL)

}

func SetUpTestDB(db *gorm.DB) {
	//Execute SQL script to create test data
	db = db.Exec(seedStudentsTeachersSQL)
}

func TearDownTestDB(db *gorm.DB) {
	//Execute SQL script to delete test data
	db = db.Exec(dropTablesSQL)
}
