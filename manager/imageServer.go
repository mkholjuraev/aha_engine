package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/publico_engine/base/models"
	"github.com/mkholjuraev/publico_engine/db/admin"
)

func ImageServer(ctx *gin.Context) {
	var image models.Images
	imageId := ctx.Param("image_id")
	if imageId == "" {
		ctx.JSON(http.StatusBadRequest, "You need to pass image id")
		return
	}

	db := admin.DB

	if err := db.Where("id = ?", imageId).First(&image).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Image not found")
		return
	}

	fmt.Printf("Image found in: %v\n", image.Path)

	ctx.File(image.Path)
}
