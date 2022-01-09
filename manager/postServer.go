package manager

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

type Post struct {
	Id          uint   `json:"id" query:"p.id"`
	Title       string `json:"title" query:"p.title"`
	Description string `json:"description" query:"p.description"`
	WriterId    uint   `json:"writer_id" query:"p.writer_id"`
	CreatedAt   string `json:"created_at" query:"created_at"`
	Name        string `json:"name" query:"u.name"`
	Surname     string `json:"surname" query:"u.surname"`
	CoverImage  string `json:"cover_image" query:"p.cover_image"`
	ReadTime    int    `json:"read_time" query:"p.read_time"`
}

//TODO: fix biography column typo
type SinglePostResponse struct {
	Post          `gorm:"embedded;embeddedPrefix:m1_"`
	Content       string `json:"content" query:"p.content"`
	Name          string `json:"name" query:"u.name"`
	Surname       string `json:"surname" query:"u.surname"`
	Biograph      string `json:"biograph" query:"w.biograph"`
	PhotoID       uint   `json:"photo_id" query:"u.photo_id"`
	Profession    string `json:"profession" query:"w.profession"`
	DistinctLikes int    `json:"distinct_likes" query:"w.distinct_likes"`
	DistinctViews int    `json:"distinct_views" query:"w.distinct_views" `
	// Specializations []string `json:"specializations" query:"specializations" `
}

type PostsFilterRequests struct {
	WriterID string `json:"writer_id"`
}

type Metadata struct {
	TotalRows int64 `json:"total_rows"`
}

type PostsResponse struct {
	Posts []Post
	Metadata
}

func PostServer(ctx *gin.Context) {
	var post SinglePostResponse

	postID := ctx.Param("post_id")
	if postID == "" {
		ctx.JSON(http.StatusBadRequest, "You need to send correct post id")
	}

	db := admin.DB

	//TODO: fix biography column typo
	err := db.Table("posts as p").
		Joins("JOIN writers as w ON w.id = p.writer_id").
		Joins("JOIN users as u ON u.id = w.user_id").
		Where("p.id = (?)", postID).
		Select("p.id as m1_id, p.title as m1_title, p.description as m1_description, p.cover_image as m1_cover_image, p.content, p.created_at::date as m1_created_at, p.writer_id as m1_writer_id, p.read_time as m1_read_time, u.name, u.surname, w.id as writer_id, u.photo_id, w.biograph, w.profession, w.distinct_likes,w.distinct_views").
		Scan(&post).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	fmt.Printf("Post served: %v\n", post.Id)
	ctx.JSON(http.StatusOK, post)
}

func PostsServer(ctx *gin.Context) {
	filters := PostsFilterRequests{
		WriterID: ctx.Request.URL.Query().Get("writer_id"),
	}

	offset, err := strconv.Atoi(ctx.Request.URL.Query().Get("offset"))
	if err != nil {
		fmt.Println("error converting offset string to int: " + err.Error())
		return
	}

	limit, err := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))
	if err != nil {
		fmt.Println("error converting limit string to int: " + err.Error())
		return
	}

	db := admin.DB
	var posts []Post
	var count int64
	dbResponse := db.Debug().Table("posts as p").
		Joins("join users as u on p.writer_id = u.id").
		Select("p.id, p.title, p.description, p.cover_image, p.created_at::date, p.writer_id, p.read_time, u.name, u.surname, u.username").
		Where(filters).
		Count(&count).
		Offset(offset).
		Limit(limit).Scan(&posts)

	if dbResponse.Error != nil {
		ctx.JSON(http.StatusBadRequest, "Posts not found")
		return
	}

	response := PostsResponse{
		Posts: posts,
		Metadata: Metadata{
			TotalRows: count,
		},
	}
	ctx.JSON(http.StatusOK, &response)
}
