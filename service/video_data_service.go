package service

import (
	"fam/models"
	"fam/storage"
)

type VideoDataService interface {
	GetData() ([]models.VideoData, error)
	SearchData(query string) ([]models.VideoData, error)
}

type videoDataService struct {
	videoDataStorage storage.VideoDataStorage
}

func NewVideoDataService(videoDataStorage storage.VideoDataStorage) VideoDataService {
	return &videoDataService{videoDataStorage: videoDataStorage}
}

func (v videoDataService) GetData() ([]models.VideoData, error) {
	var videoData []models.VideoData
	res, err := v.videoDataStorage.Get()
	if err != nil {
		return videoData, err
	}
	var datum models.VideoData
	for res.Next() {
		res.Scan(&datum.Id, &datum.Title, &datum.Description, &datum.PublishedTime, &datum.Url)
		videoData = append(videoData, datum)
	}
	return videoData, err
}

func (v videoDataService) SearchData(query string) ([]models.VideoData, error) {
	var videoData []models.VideoData
	res, err := v.videoDataStorage.Search(query)
	if err != nil {
		return videoData, err
	}
	var datum models.VideoData
	for res.Next() {
		res.Scan(&datum.Id, &datum.Title, &datum.Description, &datum.PublishedTime, &datum.Url)
		videoData = append(videoData, datum)
	}
	return videoData, err
}
