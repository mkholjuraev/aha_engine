package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/publico_engine/manager/auth"
)

const (
	authHeaderkey  = "Authorization"
	atuhTypeBearer = "bearer"
	authPayloadKey = "authorization_payload"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func AuthMiddleware(tokenMaker auth.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.Request.Header["Authorization"]

		if len(authHeader) == 0 {
			fmt.Printf("authorization header is not provided")
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(strings.Join(authHeader, " "))

		if len(fields) > 2 {
			fmt.Printf("invalid authorization header format")
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authType := strings.ToLower(fields[0])

		if authType != atuhTypeBearer {
			fmt.Printf("Unsupported authorization type %s", authType)
			err := fmt.Errorf("Unsupported authorization type %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			fmt.Printf("Authentification error: %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authPayloadKey, payload)
		ctx.Next()
	}
}
