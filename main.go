package main

import (
	"os"
	"golang-go-restaurant-management/database"
	"golang-go-restaurant-management/routes"
	"golang-go-restaurant-management/middelware"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

}
