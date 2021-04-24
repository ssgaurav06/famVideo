package router

import (
	"fam/config"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	videoDataHandler := config.Init()
	router := mux.NewRouter()
	router.HandleFunc("/videoData", videoDataHandler.GetData).Methods(http.MethodGet)
	router.HandleFunc("/search", videoDataHandler.SearchData).Methods(http.MethodGet)
	return router
}
