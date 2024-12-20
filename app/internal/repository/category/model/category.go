package model

type Category struct {
	Id             int64  `db:"id"`
	Name           string `db:"name"`
	Slug           string `db:"slug"`
	SeoTitle       string `db:"seo_title"`
	SeoDescription string `db:"seo_description"`
	PreviewPhoto   string `db:"preview_photo"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}
