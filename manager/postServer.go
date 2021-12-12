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

//Use to serve content also: p.content,   Content     string `json:"content" query:"p.content"`
type PostsResponse struct {
	Id          uint   `json:"id" query:"p.id"`
	Title       string `json:"title" query:"p.title"`
	Description string `json:"description" query:"p.description"`
	WriterId    uint   `json:"writer_id" query:"p.writer_id"`
	CreatedAt   string `json:"created_at" query:"created_at"`
	Name        string `json:"name" query:"u.name"`
	Surname     string `json:"surname" query:"u.surname"`
	Username    string `json:"username" query:"u.username"`
	CoverImage  string `json:"cover_image" query:"p.cover_image"`
	ReadTime    int    `json:"read_time" query:"p.read_time"`
}

func PostsServer(ctx *gin.Context) {

	db := admin.DB
	var response []PostsResponse

	dbResponse := db.Table("posts as p").
		Joins("join users as u on p.writer_id = u.id").
		Select("p.id, p.title, p.description, p.cover_image, p.created_at::date, p.writer_id, p.read_time, u.name, u.surname, u.username").
		Limit(10).Scan(&response)

	if dbResponse.Error != nil {
		ctx.JSON(http.StatusBadRequest, "Posts not found")
		return
	}

	ctx.JSON(http.StatusOK, &response)
}
