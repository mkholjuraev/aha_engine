package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

type UpdateRequestAttributes struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	CoverImage  string `json:"cover_image,omitempty"`
	ReadTime    int    `json:"read_time,omitempty"`
	Views       int    `json:"views,omitempty"`
	Likes       int    `json:"likes,omitempty"`
	Shares      int    `json:"shares,omitempty"`
}

func UpdatePost(ctx *gin.Context) {
	postID := ctx.Param("post_id")

	db := admin.DB

	var requestBody UpdateRequestAttributes
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var post models.Post
	post = models.Post{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Content:     requestBody.Content,
		CoverImage:  requestBody.CoverImage,
		ReadTime:    requestBody.ReadTime,
		Views:       requestBody.Views,
		Likes:       requestBody.Likes,
		Shares:      requestBody.Shares,
	}
	if err := db.Model(models.Post{}).Where("id = ?", postID).Updates(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, post.ID)
}
