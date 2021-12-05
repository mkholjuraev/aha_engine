package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

func PostServer(ctx *gin.Context) {
	var post models.Post

	postId := ctx.Param("post_id")
	if postId == "" {
		ctx.JSON(http.StatusBadRequest, "You need to send correct post id")
	}

	db := admin.DB

	if err := db.Where("id = ?", postId).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Image not found")
		return
	}

	fmt.Printf("Post served: %v\n", post.ID)
	ctx.JSON(http.StatusOK, post)
}

type PostsResponse struct {
	Id          uint   `json:"id" query:"p.id"`
	Title       string `json:"title" query:"p.title"`
	Description string `json:"description" query:"p.description"`
	Content     string `json:"content" query:"p.content"`
	WriterId    uint   `json:"writer_id" query:"p.writer_id"`
	Name        string `json:"name" query:"u.name"`
	Surname     string `json:"surname" query:"u.surname"`
	Username    string `json:"username" query:"u.username"`
}

func PostsServer(ctx *gin.Context) {

	db := admin.DB
	var response []PostsResponse

	dbResponse := db.Table("posts as p").
		Joins("join users as u on p.writer_id = u.id").
		Select("p.id, p.title, p.description, p.content, p.writer_id, u.name, u.surname, u.username").
		Limit(10).Scan(&response)

	if dbResponse.Error != nil {
		ctx.JSON(http.StatusBadRequest, "Posts not found")
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
