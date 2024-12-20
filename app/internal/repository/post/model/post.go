package model

import (
	"time"
)

type Post struct {
	Id              int64     `db:"id"`
	Title           string    `db:"title"`
	Content         string    `db:"content"`
	SeoTitle        string    `db:"seo_title"`
	SeoDescription  string    `db:"seo_description"`
	PreviewImageURL string    `db:"preview_image_url"`
	CategoryId      int64     `db:"category_id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
