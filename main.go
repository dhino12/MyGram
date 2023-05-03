package main

import (
	"fmt"
	"mygram/database"
	"mygram/routes"
)

func main() {
	database.StartDB()
	r := routes.StartApp()
	
	fmt.Println("App running on localhost:8080")
	r.Run(":8080")
}