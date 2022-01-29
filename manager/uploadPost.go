package manager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/publico_engine/base/models"
	"github.com/mkholjuraev/publico_engine/db/admin"
	"gorm.io/gorm/clause"
)

type PostRequestAttributes struct {
	Title       string   `json:"title" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	WriterID    uint     `json:"writer_id"`
	User        User     `gorm:"foreignKey:WriterID"`
	Content     string   `json:"content"`
	Views       int      `json:"views" gorm:"default:null"`
	Likes       int      `json:"likes" gorm:"default:null"`
	Shares      int      `json:"shares" gorm:"default:null"`
	CoverImage  string   `json:"cover_image"`
	ReadTime    int      `json:"read_time"`
	Theme       int      `json:"theme,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

func UploadPost(ctx *gin.Context) {

	db := admin.DB

	var requstBody PostRequestAttributes
	if err := ctx.ShouldBindJSON(&requstBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	post := mapRequestModel(requstBody)
	if err := db.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	tags, isEmpty := mapTagsModel(requstBody.Tags)
	if isEmpty != true {
		tagsQuery := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).Create(&tags)

		if tagsQuery.Error != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error in database job: %s", tagsQuery.Error.Error()))
			return
		}
	}

	metadata := mapPostMetadata(tags, requstBody, post)
	metadataQuery := db.Debug().Model(models.PostMetadata{}).Create(&metadata)

	if metadataQuery.Error != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error in database job: %s", metadataQuery.Error.Error()))
		return
	}

	ctx.JSON(http.StatusOK, post.ID)
}

func mapRequestModel(requestBody PostRequestAttributes) models.Post {
	var post models.Post
	post = models.Post{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		WriterID:    requestBody.WriterID,
		Content:     requestBody.Content,
		CoverImage:  requestBody.CoverImage,
		ReadTime:    requestBody.ReadTime,
	}

	return post
}

func mapTagsModel(tags []string) ([]models.Tags, bool) {
	if len(tags) == 0 {
		return nil, true
	}

	tagModels := make([]models.Tags, len(tags))

	for i := 0; i < len(tags); i++ {
		tagModels[i] = models.Tags{
			Name: tags[i],
		}
	}
	return tagModels, false
}

type TagIDs struct {
	IDS []uint `json:"ids"`
}

func mapPostMetadata(tags []models.Tags, requestBody PostRequestAttributes, post models.Post) models.PostMetadata {
	IDs := make([]uint, len(tags))
	for i := 0; i < len(tags); i++ {
		IDs[i] = tags[i].ID
	}

	tagIDs := TagIDs{IDS: IDs}
	tagIDJson, _ := json.Marshal(tagIDs)

	metadata := models.PostMetadata{
		TagIDJSON:        tagIDJson,
		PostID:           post.ID,
		SpecializationID: uint(requestBody.Theme),
	}

	return metadata
}
