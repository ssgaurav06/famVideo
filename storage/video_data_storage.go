package storage

import (
	"database/sql"

	"google.golang.org/api/youtube/v3"
)

type VideoDataStorage interface {
	Insert(item *youtube.SearchResult) error
	ClearDB() error
	Get() (*sql.Rows, error)
	Search(query string) (*sql.Rows, error)
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

func (vds VideoDataDB) Get() (*sql.Rows, error) {
	sqlSt := `Select * from videoData ORDER BY published_time desc`
	res, err := vds.DB.Query(sqlSt)
	return res, err
}

func (vds VideoDataDB) Search(query string) (*sql.Rows, error) {
	sqlSt := `Select * from videoData WHERE to_tsvector(title) @@ to_tsquery($1) OR to_tsvector(description) @@ to_tsquery($1) ORDER BY published_time desc`
	res, err := vds.DB.Query(sqlSt, query)
	return res, err
}
