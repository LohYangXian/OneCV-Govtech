package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"github.com/lohyangxian/OneCV-Govtech/internal/api"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: Extract DB Connection to db.go
func main() {

	// TODO: Extract Config to a function
	// Set the file name of configuration file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	// Unmarshal the configuration file into the struct
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	//TESTING
	//TODO: REMOVE ONCE DONE
	// Reading variables using the model
	fmt.Println("Reading variables using the model..")

	fmt.Println("Connected on port \t", configuration.Server.Port)
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("and Port is\t\t", configuration.Database.Port)
	fmt.Println("Environment is\t", configuration.Environment)
	fmt.Println("DB User is\t", configuration.Database.DBUser)
	fmt.Println("DB Password is\t", configuration.Database.DBPassword)

	// Open a connection to the database
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		configuration.Database.DBUser,
		configuration.Database.DBPassword,
		configuration.Database.DBName,
		configuration.Database.Port,
	)

	//TODO: Extract this to a function
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database, %v", err)
		return
	}

	server, err := api.NewServer(&configuration, db)
	if err != nil {
		fmt.Printf("Error starting server, %v", err)
		return
	}

	//TODO: Extract this
	port := string(configuration.Server.Port)
	var address string
	if port == "" {
		address = ":3000"
	} else {
		address = ":" + port
	}

	err = server.Start(address)
	if err != nil {
		fmt.Printf("Error starting server, %v", err)
		return
	}

	fmt.Println("Successfully connected to the database!")
}
