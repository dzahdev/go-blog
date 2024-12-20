package converter

import (
	"dzrise.ru/internal/model"
	modelCat "dzrise.ru/internal/repository/category/model"
	modelPost "dzrise.ru/internal/repository/post/model"
)

func ToPostFromService(post *model.Post) *modelPost.Post {
	return &modelPost.Post{
		Id:              post.Id,
		Title:           post.Title,
		Content:         post.Content,
		SeoTitle:        post.SeoTitle,
		SeoDescription:  post.SeoDescription,
		PreviewImageURL: post.PreviewImageURL,
		CategoryId:      post.Category.Id,
	}
}

func ToPostFromRepository(post *modelPost.Post, cat *modelCat.Category) *model.Post {
	return &model.Post{
		Id:              post.Id,
		Title:           post.Title,
		Content:         post.Content,
		SeoTitle:        post.SeoTitle,
		SeoDescription:  post.SeoDescription,
		PreviewImageURL: post.PreviewImageURL,
		Category: model.Category{
			Id:   cat.Id,
			Name: cat.Name,
			Slug: cat.Slug,
		},
		CreatedAt: post.CreatedAt.String(),
		UpdatedAt: post.UpdatedAt.String(),
	}
}
