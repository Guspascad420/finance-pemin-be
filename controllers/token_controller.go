package controllers

import (
	"errors"
	"finance-be/auth"
	"finance-be/database"
	"finance-be/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	record := database.Db.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}
	jwtToken, err := auth.GenerateJWT(user.Email)
	if err != nil {
		//Handle Error
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "success", "token": jwtToken})
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		c.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.Db.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": record.Error.Error()})
		c.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid Credentials"})
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "success", "data": models.LoginResponse{tokenString}})
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("unauthorized Error: Access Denied")
}

func ExtractEmail(c *gin.Context) (string, error) {
	signedToken, err := ExtractToken(c)
	if err != nil {
		return "", err
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&auth.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*auth.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}
	return claims.Email, nil
}
