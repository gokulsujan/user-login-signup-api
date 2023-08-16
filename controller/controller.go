package controller

import (
	"net/http"
	"user-registration-sinin/config"
	"user-registration-sinin/models"

	"github.com/gin-gonic/gin"
)

var user = models.User{}

func CreateUser(c *gin.Context) {
	//binding the json data to user fields
	c.ShouldBindJSON(&user)
	//Inserting into db
	result := config.DB.Create(&user)
	//if error occured
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	//if error not occured
	c.JSON(http.StatusAccepted, gin.H{"Message": "User details Inserted Successfully"})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	config.DB.First(&user, id)
	c.JSON(http.StatusAccepted, gin.H{"message": user})

}
