package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mkholjuraev/aha_engine/db/admin"
	"github.com/mkholjuraev/aha_engine/manager"
	"github.com/mkholjuraev/aha_engine/manager/auth"
	"github.com/mkholjuraev/aha_engine/manager/middleware"
)

func main() {

	admin.NewDatabaseConncetion()

	router := gin.Default()
	router.Use(middleware.CORSMiddlewarePermitLogin())
	router.POST("/api/login", auth.Login)
	router.POST("/api/register", manager.Register)
	router.GET("/api/image/:image_id", manager.ImageServer)
	router.GET("/api/posts", manager.PostsServer)
	router.GET("/api/post/:post_id", manager.PostServer)
	router.GET("/api/writers", manager.WritersListServer)
	router.GET("/api/writers/:writer_id", manager.WriterInfoServer)
	router.GET("/api/specializations", manager.SpecializationsServer)

	maker, err := auth.NewJWTMaker("xZ4PG7VtzqzHUBzDvA9EzzXiZ4nCataJ")
	authRoutes := router.Group("/").Use(middleware.CORSMiddlewarePermitAfterAuth())
	if err != nil {
		fmt.Println(err)
	}

	authRoutes.Use(middleware.AuthMiddleware(maker))
	authRoutes.GET("/api/profile", manager.Profile)
	authRoutes.POST("/api/image", manager.UploadImage)
	authRoutes.POST("/api/post/upload", manager.UploadPost)
	authRoutes.DELETE("/api/post/:post_id", manager.DeletePostHandler)
	authRoutes.POST("/api/writers", manager.UploadWriterInfo)
	authRoutes.POST("/api/specializations", manager.UploadSpecialization)
	router.Run(":8085")
}
