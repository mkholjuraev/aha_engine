package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

type WriterRequestAttributes struct {
	UserID          uint   `json:"user_id"`
	Profession      string `json:"profession"`
	Biography       string `json:"biography"`
	Specializations []int  `json:"specializations"`
}

func UploadWriterInfo(ctx *gin.Context) {

	db := admin.DB

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var requstBody WriterRequestAttributes
	err = json.Unmarshal(body, &requstBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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
		Biography:  requestBody.Biography,
	}

	return writer
}
