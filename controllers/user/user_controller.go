package user

import (
	"dnj-backend/database"
	"dnj-backend/models"
	"dnj-backend/utils"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context) {
	var data models.RegisterUserDto
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingUser := models.User{}
	err := database.DB.Where("username = ? OR email = ?", data.Username, data.Email).First(&existingUser).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Username or email already exists",
		})
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate user"})
		return
	}

	password, err := utils.HashPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var defaultRole models.Role
	if err := database.DB.First(&defaultRole, "role_name = ?", "user").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var newUser = models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: password,
		RoleID:   defaultRole.ID,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func LoginUser(c *gin.Context) {
	var data models.LoginUserDto
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if err := database.DB.Preload("Role").First(&user, "username = ? OR email = ?", data.Username, data.Username).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username",
		})
		return
	}

	if !utils.IsValidPassword(data.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}

	log.Printf("User ID: %d, Email: %s, Role: %s", user.ID, user.Email, user.Role.RoleName)

	token, err := utils.GenerateBearerToken(user.ID, user.Email, user.Role.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
