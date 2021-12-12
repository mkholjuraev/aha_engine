package manager

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

type PostRequestAttributes struct {
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	WriterID    uint   `json:"writer_id"`
	User        User   `gorm:"foreignKey:WriterID"`
	Content     string `json:"content"`
	Views       int    `json:"views" gorm:"default:null"`
	Likes       int    `json:"likes" gorm:"default:null"`
	Shares      int    `json:"shares" gorm:"default:null"`
	CoverImage  string `json:"cover_image"`
	ReadTime    string `json:"read_time"`
}

func UploadPost(ctx *gin.Context) {

	db := admin.DB

	var requstBody PostRequestAttributes
	if err := ctx.ShouldBindJSON(&requstBody); err != nil {
		ctx.JSON(http.StatusBadRequest, "You should provide correct data")
		return
	}

	post := mapRequestModel(requstBody)
	if err := db.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, post.ID)
}

func mapRequestModel(requestBody PostRequestAttributes) models.Post {
	readTime, _ := strconv.Atoi(requestBody.ReadTime)
	var post models.Post
	post = models.Post{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		WriterID:    requestBody.WriterID,
		Content:     requestBody.Content,
		CoverImage:  requestBody.CoverImage,
		ReadTime:    readTime,
	}

	return post
}
