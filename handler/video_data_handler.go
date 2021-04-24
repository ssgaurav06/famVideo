package handler

import (
	"fam/service"
	"net/http"
)

type VideoDataHandler interface {
	GetData(w http.ResponseWriter, r *http.Request)
}

type videoDataHandler struct {
	videoDataService service.VideoDataService
}

func NewVideoDataHandler(videoDataService service.VideoDataService) VideoDataHandler {
	return &videoDataHandler{videoDataService: videoDataService}
}

func (v videoDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
