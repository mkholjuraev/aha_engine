package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
)

var jwtKey = []byte("sectet")

type UserInfo struct {
	WriterId    uint `json:"writer_id" query:"writer_id"`
	models.User `gorm:"embedded;embeddedPrefix:m1_"`
}

type Credentials struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	UserInfo
	JwtToken string `json:"jwtToken"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func Login(ctx *gin.Context) {

	var credentials Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	db := admin.DB
	var userInfo UserInfo
	query := db.Table("users u").Joins("LEFT JOIN writers w ON w.user_id = u.id").
		Where("username = ? and password = ?", credentials.Username, credentials.Password).
		Select("u.id as m1_id, u.name as m1_name, u.surname as m1_surname, u.username as  m1_username, u.password as m1_password, u.photo_id as m1_photo_id, w.id as writer_id").
		Find(&userInfo)

	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, query.Error)
		return
	}

	maker, err := NewJWTMaker("xZ4PG7VtzqzHUBzDvA9EzzXiZ4nCataJ")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tokenString, err := maker.CreateToken(credentials.Username, 1)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Printf("Logged in: %v\n", userInfo.ID)
	ctx.SetCookie(fmt.Sprintf("token", userInfo.Username), tokenString, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, Response{
		UserInfo: userInfo,
		JwtToken: tokenString,
	})

}
