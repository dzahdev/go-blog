package comment

import (
	"context"
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/repository"
	model "dzrise.ru/internal/repository/comment/model"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	tableName = "comments"

	idColumn        = "id"
	postIDColumn    = "post_id"
	userIDColumn    = "user_id"
	contentColumn   = "content"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

var _ repository.CommentRepository = (*repo)(nil)

type repo struct {
	db db.Client
	tm db.TxManager
}

func New(db db.Client, tm db.TxManager) repository.CommentRepository {
	return &repo{
		db: db,
		tm: tm,
	}
}

func (r *repo) Create(ctx context.Context, comment *model.Comment) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(postIDColumn, userIDColumn, contentColumn).
		Values(comment.PostID, comment.UserID, comment.Content).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "comment_repository.Create",
		QueryRow: query,
	}

	var id int64
	if err = r.db.DB().ScanOneContext(ctx, &id, q, args...); err != nil {
		return 0, err
	}

	return id, nil

}

func (r *repo) GetByID(ctx context.Context, id int64) (*model.Comment, error) {
	builder := sq.Select(idColumn, postIDColumn, userIDColumn, contentColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "comment_repository.GetByID",
		QueryRow: query,
	}

	var comment model.Comment
	if err = r.db.DB().ScanOneContext(ctx, &comment, q, args...); err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *repo) Update(ctx context.Context, comment *model.Comment) error {
	builder := sq.Update(tableName).
		Set(postIDColumn, comment.PostID).
		Set(userIDColumn, comment.UserID).
		Set(contentColumn, comment.Content).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: comment.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "comment_repository.Update",
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
		Name:     "comment_repository.Delete",
		QueryRow: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	return err
}

func (r *repo) GetAllByPostID(ctx context.Context, postID int64) ([]*model.Comment, error) {
	builder := sq.Select(idColumn, postIDColumn, userIDColumn, contentColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{postIDColumn: postID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "comment_repository.GetAllByPostID",
		QueryRow: query,
	}
	var comments []*model.Comment

	if err = r.db.DB().ScanAllContext(ctx, &comments, q, args...); err != nil {
		return nil, err
	}

	return comments, nil
}
