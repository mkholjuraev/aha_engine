package manager

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/publico_engine/base/models"
	"github.com/mkholjuraev/publico_engine/db/admin"
	"github.com/mkholjuraev/publico_engine/utils"
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
type SinglePostQuery struct {
	Post          `gorm:"embedded;embeddedPrefix:m1_"`
	Content       string `json:"content" query:"p.content"`
	DistinctLikes int    `json:"distinct_likes" query:"w.distinct_likes"`
	DistinctViews int    `json:"distinct_views" query:"w.distinct_views" `
	Theme         string `json:"theme" query:"s.name" `
	Tags          string `json:"tags" query:"tags" `
}

type SinglePostResponse struct {
	Post SinglePostQuery `json:"post"`
	Tags []string        `json:"tags"`
}

type PostsFilterRequests struct {
	WriterID         string `json:"writer_id" query:"p.writer_id"`
	SpecializationID string `json:"specialization_id" query:"m.specialization_id"`
}

type Metadata struct {
	TotalRows int64 `json:"total_rows"`
}

type PostsResponse struct {
	Posts []Post
	Metadata
}

func PostServer(ctx *gin.Context) {
	var post SinglePostQuery

	postID := ctx.Param("post_id")
	if postID == "" {
		ctx.JSON(http.StatusBadRequest, "You need to send correct post id")
	}

	db := admin.DB

	//TODO: fix biography column typo
	err := db.Table("posts as p").
		Joins("JOIN writers as w ON w.id = p.writer_id").
		Joins("JOIN users as u ON u.id = w.user_id").
		Joins("LEFT JOIN post_metadata as m ON m.post_id = p.id").
		Joins("LEFT JOIN specializations as s ON s.id = m.specialization_id").
		Where("p.id = (?)", postID).
		Select(`p.id as m1_id, p.title as m1_title, p.description as m1_description, p.cover_image as m1_cover_image, p.content, 
			p.created_at::date as m1_created_at, p.writer_id as m1_writer_id, p.read_time as m1_read_time, w.id as writer_id,
			w.distinct_likes,w.distinct_views, s.name as theme, m.tag_id_json as tags`).Scan(&post).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var tagIDs TagIDs

	var tagNames []string
	if post.Tags != "" {
		if err := json.Unmarshal([]byte(post.Tags), &tagIDs); err != nil {
			panic(err)
		}

		err = db.Model(models.Tags{}).Select("name").Where("id in (?)", tagIDs.IDS).Scan(&tagNames).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	response := SinglePostResponse{
		Post: post,
		Tags: tagNames,
	}

	ctx.JSON(http.StatusOK, response)
}

func PostsServer(ctx *gin.Context) {
	filters := PostsFilterRequests{
		WriterID:         ctx.Request.URL.Query().Get("writer_id"),
		SpecializationID: ctx.Request.URL.Query().Get("specialization_id"),
	}

	db := admin.DB
	var posts []Post
	var count int64
	allRows := db.Debug().Table("posts as p").
		Joins("JOIN writers AS w ON p.writer_id = w.id").
		Joins("JOIN users AS u ON u.id = w.user_id").
		Joins("LEFT JOIN post_metadata AS m ON m.post_id = p.id").
		Select("p.id, p.title, p.description, p.cover_image, p.created_at::date, p.writer_id, p.read_time, u.name, u.surname, u.username")

	whereCondition := utils.QueryConditions(filters)
	allRows.Where(strings.Join(whereCondition, " AND "))

	offsetString := ctx.Request.URL.Query().Get("offset")
	if offsetString != "" {
		offset, _ := strconv.Atoi(offsetString)
		allRows.Count(&count).Offset(offset)
	}

	limitString := ctx.Request.URL.Query().Get("limit")
	if limitString != "" {
		limit, _ := strconv.Atoi(limitString)
		allRows.Limit(limit)
	}

	finalQuery := allRows.Scan(&posts)

	if finalQuery.Error != nil {
		ctx.JSON(http.StatusBadRequest, finalQuery.Error.Error())
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
