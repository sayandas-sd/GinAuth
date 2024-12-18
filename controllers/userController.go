package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sayandas-sd/GinAuth/database"
	"github.com/sayandas-sd/GinAuth/models"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var input models.User
		var existingUser models.User

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := validate.Struct(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		if err := database.DB.Where("email=?", input.Email).First(&existingUser); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Email is already taken",
			})
		}

		hashsedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error hashing password",
			})
		}

		newUser := models.User{
			ID:        uuid.New(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  string(hashsedPassword),
			Phone:     input.Phone,
			User_Type: input.User_Type,
		}

		if err := database.DB.Create(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create user",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "user data create successfully",
			"newuser": newUser,
		})

	}
}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		if err := helper.MatchUserType(c, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		var user []models.User

		db := database.DB

		if err := db.First(&user, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})

	}
}

func HashPassword() {

}
