package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

type WritersListResponse struct {
	UserId   uint   `json:"user_id" query:"u.id"`
	Name     string `json:"name" query:"u.name"`
	Surname  string `json:"surname" query:"u.surname"`
	WriterId uint   `json:"writer_id" query:"w.writer_id"`
	PhotoID  uint   `json:"photo_id" query:"u.photo_id"`
}

type WriterResponse struct {
	Biography     string `json:"biography" query:"w.biography"`
	Profession    string `json:"profession" query:"w.profession"`
	DistinctLikes int    `json:"distinct_likes" query:"w.distinct_likes"`
	DistinctViews int    `json:"distinct_views" query:"w.distinct_views"`
}

func WritersListServer(ctx *gin.Context) {

	db := admin.DB
	var response []WritersListResponse

	dbResponse := db.Table("writers as w").
		Joins("JOIN users as u ON u.id = w.user_id").
		Select("u.id, u.name, u.surname, w.id as writer_id, u.photo_id").
		Limit(10).Scan(&response)

	if dbResponse.Error != nil {
		ctx.JSON(http.StatusBadRequest, dbResponse.Error)
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func WriterInfoServer(ctx *gin.Context) {

	db := admin.DB
	var response []WritersListResponse
	writerID := ctx.Param("writer_id")

	dbResponse := db.Table("writers as w").
		Joins("JOIN users as u ON u.id = w.user_id").
		Select("u.id, u.name, u.surname, w.id as writer_id, u.photo_id, w.biography, w.profession, w.distinct_likes, w.distinct_views").
		Where("w.id = (?)", writerID).
		Limit(10).Scan(&response)

	if dbResponse.Error != nil {
		ctx.JSON(http.StatusBadRequest, dbResponse.Error)
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
