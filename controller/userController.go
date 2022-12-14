package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	initializer "github.com/techlateef/go-auth/Initializer"
	models "github.com/techlateef/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	//Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return

	}
	//Create the User
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializer.ConnectToDb().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{})

}
