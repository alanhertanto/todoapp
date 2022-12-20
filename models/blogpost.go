package models

import (
	"time"
)

type Post struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Source      string    `json:"source"`
	Category    string    `json:"category"`
	Language    string    `json:"language"`
	Country     string    `json:"country"`
	PublishedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"published_at"`
}
