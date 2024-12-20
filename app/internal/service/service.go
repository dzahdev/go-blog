package service

import (
	"context"
	"dzrise.ru/internal/model"
)

type PostService interface {
	Create(ctx context.Context, post *model.Post) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.Post, error)
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*model.Post, error)
	GetAllByCategoryId(ctx context.Context, categoryID int64) ([]*model.Post, error)
}

type CategoryService interface {
	Create(ctx context.Context, category *model.Category) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.Category, error)
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetBySlug(ctx context.Context, slug string) (*model.Category, error)
}
