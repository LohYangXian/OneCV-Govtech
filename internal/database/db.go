package database

import (
	"fmt"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(configuration config.Configurations) *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
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

	return db
}
