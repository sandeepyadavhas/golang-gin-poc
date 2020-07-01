package service

import (
	"github.com/sandeepyadavhas/golang-gin-poc/entity"
)

// VideoService abc
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// New abc
func New() VideoService {
	vs := videoService{}
	vs.videos = make([]entity.Video, 0)
	return &vs
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
