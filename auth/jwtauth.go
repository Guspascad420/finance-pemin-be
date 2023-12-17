package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (tokenString string, err error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
