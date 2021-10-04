package main

import (
	"fmt"
	"log"

	config "delivery-validation/config"
	"delivery-validation/pkg/database"
	"delivery-validation/pkg/handlers"
	logger "delivery-validation/pkg/logger"
	"delivery-validation/pkg/router"
)

func main() {
	logger := logger.NewLogger("log.txt")
	logger.InfoLogger.Println("Reading database configuration")

	/*
		databaseConfig, err := config.LoadDatabaseConfiguration()
		if err != nil {
			log.Printf("Error setting database : %s\n", err.Error())
			return
		}
	*/

	//initializing db and router
	logger.InfoLogger.Println("Initializing Program")

	/*
		Database, err := database.NewDatabase("mysql",
			databaseConfig.Username, databaseConfig.Password, databaseConfig.Address,
			databaseConfig.DatabaseName)
	*/

	database, err := database.NewDatabase("mysql",
		"root", "123jonathan123100300!!!", "localhost:3306",
		"testers")

	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}

	router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(database, router, logger)
	handlers.RegisterAllHandlers()
	port := config.LoadPort()
	fmt.Println("Server starting at "+ port)
	router.Start(port)
}
