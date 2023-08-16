package main

import (
	"user-registration-sinin/config"
	"user-registration-sinin/models"
)

func init() {
	config.PortInitializer()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
}
