package controller

import (
	"net/http"
	"user-registration-sinin/config"
	"user-registration-sinin/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	//binding the json data to user fields
	c.ShouldBindJSON(&user)
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"Message": "User details Inserted Successfully"})
}
