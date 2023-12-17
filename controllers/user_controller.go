package controllers

import (
	"finance-be/database"
	"finance-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserProfile(c *gin.Context) {
	var user models.User
	email, err := ExtractEmail(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": err.Error()})
		c.Abort()
		return
	}
	record := database.Db.Where("email = ?", email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true, "message": "success",
		"data": models.UserProfileResponse{user.ID, user.Name, user.Email},
	})
}
