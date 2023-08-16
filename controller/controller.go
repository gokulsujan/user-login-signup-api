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
	// if the cookie already exists redirect to profile
	if _, err := c.Request.Cookie("Email"); err == nil {
		c.Redirect(http.StatusSeeOther, "/profile")
	} else { // if not exists verify credentials show the message

		//getting the details from the form
		FormEmail := c.PostForm("email")
		FormPassword := c.PostForm("password")

		//checking the form data with database
		result := config.DB.Where(&models.User{Email: FormEmail, Password: FormPassword}).Find(&user)
		// if the datas not found show invalid credentials
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid credentials"})
		} else { // if the data found create the cookie with email
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

}

func GetUser(c *gin.Context) {
	// getting the user details with cookie
	if Cookie, err := c.Request.Cookie("Email"); err == nil {
		id := Cookie.Value
		config.DB.Where("email = ?", id).First(&user)
		c.JSON(http.StatusAccepted, gin.H{"message": user})

	} else {
		c.Redirect(http.StatusSeeOther, "/login")
	}

}

func UpdateUser(c *gin.Context) {
	var body struct {
		Name     string `json:"Name"`
		Age      int    `json:"Age"`
		Gender   string `json:"Gender"`
		Mobile   string `json:"Mobile"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	c.ShouldBindJSON(&body)

	Cookie, _ := c.Request.Cookie("Email")
	config.DB.Where("email = ?", Cookie.Value).First(&user)
	result := config.DB.Model(&user).Updates(models.User{Name: body.Name, Age: body.Age, Gender: body.Gender, Mobile: body.Mobile, Email: body.Email, Password: body.Password})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Updated"})

}
