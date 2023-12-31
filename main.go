package main

import (
	"user-registration-sinin/config"
	"user-registration-sinin/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	config.PortInitializer()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hai"})
	})
	r.POST("/create_user", controller.CreateUser)
	r.GET("/profile", controller.GetUser)
	r.POST("/login", controller.VerifyUser)
	r.POST("/update", controller.UpdateUser)

	r.Run()
}
