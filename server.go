package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/polonkoevv/goGin/controller"
	"github.com/polonkoevv/goGin/middlewares"
	"github.com/polonkoevv/goGin/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func SetupLogFile() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {

	SetupLogFile()

	server := gin.New()

	server.Use(middlewares.Logger(), gin.Recovery(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {

			ctx.JSON(http.StatusOK, gin.H{"message": "Video message is valid"})
		}

	})

	server.Run(":8080")
}
