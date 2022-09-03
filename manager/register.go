package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/publico_engine/base/models"
	"github.com/mkholjuraev/publico_engine/db/admin"
	"github.com/mkholjuraev/publico_engine/utils"
)

type User struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password" gorm:"not null"`
	Telephone   string `json:"telephone" gorm:"not null"`
	SocialLinks string `json:"social_links"`
}

func Register(ctx *gin.Context) {

	file, err := ctx.FormFile("photo")

	var profileImage models.Images
	if err == nil {
		profileImage, err = utils.SaveFile(ctx, file, "./store/images/profilePhotos/")

	}

	db := admin.DB
	response := db.Create(&profileImage)
	if response.Error != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", response.Error))
		return
	}

	var user User
	ctx.Request.ParseMultipartForm(0)

	hashedPassword, err := utils.HashPassord(ctx.Request.FormValue("password"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}

	var newUser = models.User{
		Name:        ctx.Request.FormValue("name"),
		Surname:     ctx.Request.FormValue("surname"),
		Username:    ctx.Request.FormValue("username"),
		Password:    hashedPassword,
		Telephone:   ctx.Request.FormValue("telephone"),
		SocialLinks: ctx.Request.FormValue("social_links"),
		PhotoID:     profileImage.Id,
	}
	if err := db.Create(&newUser).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, user.Name)
}
