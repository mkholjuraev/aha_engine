package manager

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mkholjuraev/aha_engine/base/models"
	"github.com/mkholjuraev/aha_engine/db/admin"
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

	// The file cannot be received.
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, "./store/images/profilePhotos/"+newFileName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	newfile := models.Images{
		Name:      file.Filename,
		Path:      "./store/images/profilePhotos/" + newFileName,
		Extension: extension,
		Type:      1,
	}

	db := admin.DB
	response := db.Create(&newfile)
	if response.Error != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", response.Error))
		return
	}

	var user User
	ctx.Request.ParseMultipartForm(0)

	var newUser = models.User{
		Name:        ctx.Request.FormValue("name"),
		Surname:     ctx.Request.FormValue("surname"),
		Username:    ctx.Request.FormValue("username"),
		Password:    ctx.Request.FormValue("password"),
		Telephone:   ctx.Request.FormValue("telephone"),
		SocialLinks: ctx.Request.FormValue("social_links"),
		PhotoID:     newfile.Id,
	}
	if err := db.Create(&newUser).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, user.Name)
}
