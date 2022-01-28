package manager

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/publico_engine/base/models"
	"github.com/mkholjuraev/publico_engine/db/admin"
)

type SpecizalizationAttributes struct {
	ID          uint   `json:"id"`
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

	var requstBody SpecizalizationAttributes
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

func mapUploadSpecizalizationRequestModel(requestBody SpecizalizationAttributes) models.Specialization {

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

type SpecizalizationResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func SpecializationsServer(ctx *gin.Context) {

	db := admin.DB
	var response []SpecizalizationResponse

	dbResponse := db.Table("specializations as s").
		Select("s.id, s.name, s.description").
		Limit(10).Scan(&response)

	if dbResponse.Error != nil {
		ctx.JSON(http.StatusBadRequest, "Posts not found")
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
