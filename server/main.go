package main

import (
	"fmt"

	"github.com/hyfic/snorlax/api/database"
	"github.com/hyfic/snorlax/api/router"
	"github.com/hyfic/snorlax/api/util"
)

func main() {
	fmt.Println("SNORLAX SERVER v1.0.0 🚀")
	fmt.Println("=======================")

	database.InitDB()         // Initialize the database
	defer database.DB.Close() // Ensure the database connection is closed when the program exits

	// get storage folder
	util.GetStorageFolder()
	fmt.Println("=======================")

	// ask custom port
	port := util.GetPort()
	router.StartServer(port)
}
