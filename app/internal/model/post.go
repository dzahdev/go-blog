package model

type Post struct {
	Id              int64    `json:"id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	SeoTitle        string   `json:"seo_title"`
	SeoDescription  string   `json:"seo_description"`
	PreviewImageURL string   `json:"preview_image_url"`
	Category        Category `json:"category"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}
