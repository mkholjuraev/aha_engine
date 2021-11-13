package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(ctx *gin.Context) {

	fmt.Printf("request response")
	ctx.JSON(http.StatusOK, fmt.Sprintf("hello, %s", "claims.Username"))
}
