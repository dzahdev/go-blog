package post

import (
	"context"
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/repository"
	modelPost "dzrise.ru/internal/repository/post/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "posts"

	idColumn              = "id"
	categoryIDColumn      = "category_id"
	titleColumn           = "title"
	contentColumn         = "content"
	seoTitleColumn        = "seo_title"
	seoDescriptionColumn  = "seo_description"
	PreviewImageURLColumn = "preview_image_url"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

var _ repository.PostRepository = (*repo)(nil)

type repo struct {
	db db.Client
	tm db.TxManager
}

func New(db db.Client, tm db.TxManager) repository.PostRepository {
	return &repo{
		db: db,
		tm: tm,
	}
}

func (r *repo) Create(ctx context.Context, post *modelPost.Post) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(categoryIDColumn, titleColumn, contentColumn, seoTitleColumn, seoDescriptionColumn, PreviewImageURLColumn).
		Values(post.CategoryId, post.Title, post.Content, post.SeoTitle, post.SeoDescription, post.PreviewImageURL).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	q := db.Query{
		Name:     "post_repository.Create",
		QueryRow: query,
	}

	var id int64

	if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetByID(ctx context.Context, id int64) (*modelPost.Post, error) {
	builder := sq.Select(idColumn, categoryIDColumn, titleColumn, contentColumn, seoTitleColumn, seoDescriptionColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "post_repository.GetByID",
		QueryRow: query,
	}

	var post modelPost.Post
	if err = r.db.DB().ScanOneContext(ctx, &post, q, args...); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *repo) Update(ctx context.Context, post *modelPost.Post) error {
	builder := sq.Update(tableName).
		Set(categoryIDColumn, post.CategoryId).
		Set(titleColumn, post.Title).
		Set(contentColumn, post.Content).
		Set(seoTitleColumn, post.SeoTitle).
		Set(seoDescriptionColumn, post.SeoDescription).
		Where(sq.Eq{idColumn: post.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "post_repository.Update",
		QueryRow: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "post_repository.Delete",
		QueryRow: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetAll(ctx context.Context) ([]*modelPost.Post, error) {
	builder := sq.Select(idColumn, categoryIDColumn, titleColumn, contentColumn, seoTitleColumn, seoDescriptionColumn, createdAtColumn, updatedAtColumn).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "post_repository.GetAll",
		QueryRow: query,
	}
	var posts []*modelPost.Post
	err = r.db.DB().ScanAllContext(ctx, posts, q, args...)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *repo) GetAllByCategoryId(ctx context.Context, categoryID int64) ([]*modelPost.Post, error) {
	builder := sq.Select(idColumn, categoryIDColumn, titleColumn, contentColumn, seoTitleColumn, seoDescriptionColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{categoryIDColumn: categoryID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "post_repository.GetAllByCategoryId",
		QueryRow: query,
	}

	var posts []*modelPost.Post
	err = r.db.DB().ScanAllContext(ctx, &posts, q, args...)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
