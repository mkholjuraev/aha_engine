package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

func UploadPost(ctx *gin.Context) {

	db := admin.DB

	var post models.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, "You should provide correct data")
		return
	}

	if err := db.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, post.ID)
}
