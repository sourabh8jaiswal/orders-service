package main

import (
	"log"
	"order-service/initializers"
	"order-service/models"
)

func main() {

	initializers.ConnectToDB()

	err := initializers.DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	log.Println("Migration done successfully")
}
