package manager

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Profile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Printf("Unauthorized attempt: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Printf("Bad request: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	tokenString := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Printf("Unauthorized attempt: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
		}
		fmt.Printf("Bad request: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		fmt.Printf("Unauthorized attempt, token expired")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Printf("request response")
	w.Write([]byte(fmt.Sprintf("hello, %s", claims.Username)))
}
