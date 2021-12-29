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

type SpecizalizationRequestAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	WriterIDs   []uint `json:"writer_ids"`
}

func UploadSpecialization(ctx *gin.Context) {

	db := admin.DB

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var requstBody SpecizalizationRequestAttributes
	err = json.Unmarshal(body, &requstBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	specialization := mapUploadSpecizalizationRequestModel(requstBody)

	if err := db.Omit("Writers.*").Create(&specialization).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func mapUploadSpecizalizationRequestModel(requestBody SpecizalizationRequestAttributes) models.Specialization {
	fmt.Println(requestBody.WriterIDs)
	writers := make([]models.Writer, len(requestBody.WriterIDs))
	for i := 0; i < len(requestBody.WriterIDs); i++ {
		writerID := (requestBody.WriterIDs)[i]
		writers[i] = models.Writer{
			BaseModel: models.BaseModel{ID: writerID},
		}
	}
	var specialization models.Specialization
	specialization = models.Specialization{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		Writers:     writers,
	}

	return specialization
}
