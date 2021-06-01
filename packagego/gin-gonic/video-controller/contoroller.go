package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-gonic/entity"
	"gilab.com/pragmaticreviews/golang-gin-gonic/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

// func New(service service.VideoService) VideoController {
// 	return controller{
// 		service: service,
// 	}
// }

func (c *controller) FindAll() entity.Video {
	return c.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)

	return video
}
