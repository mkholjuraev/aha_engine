package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("sectet")

var users = map[string]string{
	"men":   "men",
	"user2": "password2",
}

type Credentials struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func Login(ctx *gin.Context) {

	var credentials Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	expectdPassword, ok := users[credentials.Username]

	if !ok || expectdPassword != credentials.Password {
		ctx.JSON(http.StatusUnauthorized, "")
		return
	}

	maker, err := NewJWTMaker("xZ4PG7VtzqzHUBzDvA9EzzXiZ4nCataJ")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tokenString, err := maker.CreateToken(credentials.Username, 1)
	fmt.Printf("Login attempt: %s\n", tokenString)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Printf("Logged in: %v\n", credentials.Username)

	ctx.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, tokenString)

}
