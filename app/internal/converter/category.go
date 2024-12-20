package converter

import (
	"dzrise.ru/internal/model"
	modelRepo "dzrise.ru/internal/repository/category/model"
	"time"
)

func ToCategoryFromService(c model.Category) modelRepo.Category {
	return modelRepo.Category{
		Name:           c.Name,
		Slug:           c.Slug,
		PreviewPhoto:   c.PreviewPhoto,
		SeoTitle:       c.SeoTitle,
		SeoDescription: c.SeoDescription,
		CreatedAt:      time.Now().String(),
		UpdatedAt:      time.Now().String(),
	}
}
