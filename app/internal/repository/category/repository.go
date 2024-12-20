package category

import (
	"context"
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/repository"
	model "dzrise.ru/internal/repository/category/model"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	tableName = "categories"

	idColumn             = "id"
	nameColumn           = "name"
	slugColumn           = "slug"
	seoTitleColumn       = "seo_title"
	seoDescriptionColumn = "seo_description"
	previewPhotoColumn   = "preview_photo"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
)

var _ repository.CategoryRepository = (*repo)(nil)

type repo struct {
	db db.Client
	tm db.TxManager
}

func New(db db.Client, tm db.TxManager) repository.CategoryRepository {
	return &repo{
		db: db,
		tm: tm,
	}
}

func (r *repo) Create(ctx context.Context, category *model.Category) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(nameColumn, slugColumn, seoTitleColumn, seoDescriptionColumn, previewPhotoColumn).
		Values(category.Name, category.Slug, category.SeoTitle, category.SeoDescription, category.PreviewPhoto).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "category_repository.Create",
		QueryRow: query,
	}

	var id int64
	if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, seoTitleColumn, seoDescriptionColumn, previewPhotoColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "category_repository.GetByID",
		QueryRow: query,
	}

	var category model.Category
	if err = r.db.DB().ScanOneContext(ctx, &category, q, args...); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repo) Update(ctx context.Context, category *model.Category) error {
	builder := sq.Update(tableName).
		Set(nameColumn, category.Name).
		Set(slugColumn, category.Slug).
		Set(seoTitleColumn, category.SeoTitle).
		Set(seoDescriptionColumn, category.SeoDescription).
		Set(previewPhotoColumn, category.PreviewPhoto).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: category.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "category_repository.Update",
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
		Name:     "category_repository.Delete",
		QueryRow: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetAll(ctx context.Context) ([]*model.Category, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, seoTitleColumn, seoDescriptionColumn, previewPhotoColumn, createdAtColumn, updatedAtColumn).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "category_repository.GetAll",
		QueryRow: query,
	}
	var categories []*model.Category

	if err = r.db.DB().ScanAllContext(ctx, &categories, q, args...); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repo) GetBySlug(ctx context.Context, slug string) (*model.Category, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, seoTitleColumn, seoDescriptionColumn, previewPhotoColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{slugColumn: slug})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "category_repository.GetBySlug",
		QueryRow: query,
	}

	var category model.Category
	if err = r.db.DB().ScanOneContext(ctx, &category, q, args...); err != nil {
		return nil, err
	}

	return &category, nil
}
