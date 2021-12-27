package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

type WriterRequestAttributes struct {
	UserID          uint   `json:"user_id"`
	Profession      string `json:"profession"`
	Biography       string `json:"biography"`
	Specializations []uint `json:"specializations"`
}

func UploadWriterInfo(ctx *gin.Context) {

	db := admin.DB

	var requstBody WriterRequestAttributes
	if err := ctx.ShouldBindJSON(&requstBody); err != nil {
		ctx.JSON(http.StatusBadRequest, "You should provide correct data")
		return
	}

	writer := mapUploadWriterInfoRequestModel(requstBody)
	if err := db.Create(&writer).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, writer.ID)
}

func mapUploadWriterInfoRequestModel(requestBody WriterRequestAttributes) models.Writer {
	var writer models.Writer
	writer = models.Writer{
		UserID:     requestBody.UserID,
		Profession: requestBody.Profession,
		Biograph:   requestBody.Biography,
	}

	return writer
}
