package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mkholjuraev/publico_engine/base/models"
)

func SaveFile(ctx *gin.Context, file *multipart.FileHeader, folderName string) (newFile models.Images, err error) {
	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, folderName+newFileName); err != nil {
		return newFile, err
	}

	newfile := models.Images{
		Name:      file.Filename,
		Path:      folderName + newFileName,
		Extension: extension,
		Type:      1,
	}
	return newfile, nil
}
