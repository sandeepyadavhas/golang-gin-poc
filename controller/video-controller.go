package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sandeepyadavhas/golang-gin-poc/entity"
	"github.com/sandeepyadavhas/golang-gin-poc/service"
)

// VideoController abc
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	videoService service.VideoService
}

// New abc
func New(videoService service.VideoService) VideoController {
	return &videoController{videoService: videoService}
}

func (c *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return c.videoService.Save(video)
}

func (c *videoController) FindAll() []entity.Video {
	return c.videoService.FindAll()
}
