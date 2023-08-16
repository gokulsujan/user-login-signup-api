package controller

import (
	"errors"
	"net/http"
	"user-registration-sinin/config"
	"user-registration-sinin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func VerifyUser(c *gin.Context) {
	if _, err := c.Request.Cookie("Email"); err == nil {
		c.Redirect(http.StatusSeeOther, "/profile")
	} else {
		FormEmail := c.PostForm("email")
		FormPassword := c.PostForm("password")

		result := config.DB.Where(&models.User{Email: FormEmail, Password: FormPassword}).Find(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid credentials"})
		}
		cookie := http.Cookie{
			Name:     "Email",
			Value:    user.Email,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, &cookie)
		c.JSON(http.StatusAccepted, gin.H{"message": "User logged in"})
	}

}

func GetUser(c *gin.Context) {
	if Cookie, err := c.Request.Cookie("Email"); err == nil {
		id := Cookie.Value
		config.DB.Where("email = ?", id).First(&user)
		c.JSON(http.StatusAccepted, gin.H{"message": user})

	} else {
		c.Redirect(http.StatusSeeOther, "/login")
	}

}
