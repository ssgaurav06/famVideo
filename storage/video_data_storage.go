package storage

import (
	"database/sql"

	"google.golang.org/api/youtube/v3"
)

type VideoDataStorage interface {
	Insert(item *youtube.SearchResult) error
	ClearDB() error
}

type VideoDataDB struct {
	DB *sql.DB
}

func NewVideoDataStorage(db *sql.DB) VideoDataStorage {
	return &VideoDataDB{DB: db}
}

func (vds VideoDataDB) ClearDB() error {
	sqlSt := `Delete from videoData`
	_, err := vds.DB.Exec(sqlSt)
	return err
}

func (vds *VideoDataDB) Insert(item *youtube.SearchResult) error {
	sqlSt := `Insert into videoData(id,title,description,published_time,url) values ($1,$2,$3,$4,$5)`
	_, err := vds.DB.Exec(sqlSt, item.Id.VideoId, item.Snippet.Title, item.Snippet.Description, item.Snippet.PublishedAt, item.Snippet.Thumbnails.Default.Url)
	return err
}
