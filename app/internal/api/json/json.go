package json

import "dzrise.ru/internal/service"

type JSONHandlers struct {
	postService     service.PostService
	categoryService service.CategoryService
}

func NewJSONHandlers(postService service.PostService, categoryService service.CategoryService) *JSONHandlers {
	return &JSONHandlers{
		postService:     postService,
		categoryService: categoryService,
	}
}
