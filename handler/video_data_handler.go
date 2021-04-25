package handler

import (
	"encoding/json"
	"fam/constants"
	"fam/service"
	"net/http"
)

type VideoDataHandler interface {
	GetData(w http.ResponseWriter, r *http.Request)
	SearchData(w http.ResponseWriter, r *http.Request)
}

type videoDataHandler struct {
	videoDataService service.VideoDataService
}

func NewVideoDataHandler(videoDataService service.VideoDataService) VideoDataHandler {
	return &videoDataHandler{videoDataService: videoDataService}
}

func (v videoDataHandler) GetData(w http.ResponseWriter, _ *http.Request) {
	data, err := v.videoDataService.GetData()
	w.Header().Set(constants.ContentType, constants.ApplicationJSONFormat)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorMessage(w, constants.VideoDataNotPresent)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (v videoDataHandler) SearchData(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := params[constants.Query][0]
	data, err := v.videoDataService.SearchData(query)
	w.Header().Set(constants.ContentType, constants.ApplicationJSONFormat)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorMessage(w, constants.VideoDataNotPresent)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func errorMessage(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(map[string]string{
		constants.DisplayMessage: message,
	})
}
