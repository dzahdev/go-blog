package repository

import (
	"context"
	modelCategory "dzrise.ru/internal/repository/category/model"
	modelComment "dzrise.ru/internal/repository/comment/model"
	modelPost "dzrise.ru/internal/repository/post/model"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *modelComment.Comment) (int64, error)
	GetByID(ctx context.Context, id int64) (*modelComment.Comment, error)
	Update(ctx context.Context, comment *modelComment.Comment) error
	Delete(ctx context.Context, id int64) error
	GetAllByPostID(ctx context.Context, postID int64) ([]*modelComment.Comment, error)
}

type PostRepository interface {
	Create(ctx context.Context, post *modelPost.Post) (int64, error)
	GetByID(ctx context.Context, id int64) (*modelPost.Post, error)
	Update(ctx context.Context, post *modelPost.Post) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*modelPost.Post, error)
	GetAllByCategoryId(ctx context.Context, categoryID int64) ([]*modelPost.Post, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, category *modelCategory.Category) (int64, error)
	GetByID(ctx context.Context, id int64) (*modelCategory.Category, error)
	GetBySlug(ctx context.Context, slug string) (*modelCategory.Category, error)
	Update(ctx context.Context, category *modelCategory.Category) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*modelCategory.Category, error)
}
