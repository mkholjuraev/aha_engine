package manager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	db := admin.DB

	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "You should provide correct data")
		return
	}

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, user.Name)
}
