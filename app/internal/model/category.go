package model

type Category struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	SeoTitle       string `json:"seo_title"`
	SeoDescription string `json:"seo_description"`
	PreviewPhoto   string `json:"preview_photo"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
