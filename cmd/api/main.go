package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lohyangxian/OneCV-Govtech/config"
	"github.com/lohyangxian/OneCV-Govtech/internal/api"
	"github.com/spf13/viper"
)

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

	// Construct db connection string
	connectionString := fmt.Sprintf("host=localhost port=%d dbname=%s user=%s password=%s sslmode=disable",
		configuration.Database.Port,
		configuration.Database.DBName,
		configuration.Database.DBUser,
		configuration.Database.DBPassword)

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
	//TODO: Extract this to a function
	db, err := sql.Open("postgres", connectionString)
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

	// Defer the close till after the main function has finished
	defer db.Close()

	// Test the connection
	//TODO: REMOVE ONCE DONE
	//TODO: Extract this to a function
	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database: %s\n", err)
		return
	}

	fmt.Println("Successfully connected to the database!")
}
