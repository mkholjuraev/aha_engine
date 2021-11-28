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

func UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	// The file cannot be received.
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, "./store/images/"+newFileName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	newfile := models.Images{
		Name:      file.Filename,
		Path:      "./store/images/" + newFileName,
		Extension: extension,
		Type:      1,
	}

	db := admin.DB
	response := db.Create(&newfile)
	if response.Error != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error in database job: %s", response.Error))
		return
	}

	ctx.JSON(http.StatusOK, newfile.Id)
}
