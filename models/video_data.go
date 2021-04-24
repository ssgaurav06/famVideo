package models

import "time"

type VideoData struct {
	Id            string    `json:"id" db:"id"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	PublishedTime time.Time `json:"published_time" db:"published_time"`
	Url           string    `json:"url" db:"url"`
}
