package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

func DeletePostHandler(ctx *gin.Context) {

	postId := ctx.Param("post_id")
	if postId == "" {
		ctx.JSON(http.StatusBadRequest, "Post id is not provided")
	}

	db := admin.DB

	if err := db.Delete(&models.Post{}, postId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	fmt.Printf("Post found in: %v\n", postId)

	ctx.JSON(http.StatusOK, postId)
}
