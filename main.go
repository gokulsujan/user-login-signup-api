package main

import (
	"user-registration-sinin/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.PortInitializer()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hai"})
	})

	r.Run()
}
