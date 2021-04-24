package service

import (
	"fam/models"
	"fam/storage"
)

type VideoDataService interface {
	GetData() ([]models.VideoData, error)
}

type videoDataService struct {
	videoDataStorage storage.VideoDataStorage
}

func NewVideoDataService(videoDataStorage storage.VideoDataStorage) VideoDataService {
	return &videoDataService{videoDataStorage: videoDataStorage}
}

func (v videoDataService) GetData() ([]models.VideoData, error) {
	panic("implement me")
}
