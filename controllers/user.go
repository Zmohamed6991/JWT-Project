package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Zmohamed6991/JWT-Project/config"
	"github.com/Zmohamed6991/JWT-Project/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	//GET REQUEST AND BIND
	var newUser models.User
	err := c.ShouldBind(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	// HASH PASSWORD USING BCRYPT
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
	}

	//CONNECT AND CREATE USER WITH HASH PASSWORD
	//CONVERT THE BYTE TO STRING
	user := models.User{Email: newUser.Email, Password: string(hash)}

	// ADD/CREATE IN DB
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not created in database",
		})
		return
	}

	//RESPOND WITH SUCCESS MESSAGE
	c.JSON(http.StatusCreated, gin.H{"success": "user created"})

}

func LoginUser(c *gin.Context) {

	// GET request of request body
	var userLogin models.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "email = ?", userLogin.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error generating access token",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": tokenString,
	})
}
