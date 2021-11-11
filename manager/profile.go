package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Profile(ctx *gin.Context) {

	fmt.Printf("request response")
	ctx.JSON(http.StatusOK, []byte(fmt.Sprintf("hello, %s", "claims.Username")))
}
