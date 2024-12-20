package transaction

import (
	"context"
	"dzrise.ru/internal/client/db"
	"dzrise.ru/internal/client/db/pg"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type key string

type manager struct {
	db  db.Transactor
	key key
}

func NewManager(db db.Transactor, key key) db.TxManager {
	return &manager{
		db:  db,
		key: key,
	}
}

// transaction is a helper function that executes a transaction.
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, f db.Handler) (err error) {
	_, ok := ctx.Value(m.key).(pgx.Tx)
	if ok {
		return f(ctx)
	}

	tx, err := m.db.BeginTx(ctx, opts)

	if err != nil {
		err = errors.Wrap(err, "failed to begin transaction")
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		// recover from panic
		if r := recover(); r != nil {
			err = errors.Wrapf(err, "panic recover: %v", r)
		}

		// check if we need to rollback the transaction
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "failed to rollback transaction: %v", errRollback)
			}

			return
		}
		// commit the transaction
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "failed to commit transaction")
			}
		}
	}()

	if err = f(ctx); err != nil {
		err = errors.Wrap(err, "failed to execute transaction handler")
	}

	return err

}

func (m *manager) ReadCommited(ctx context.Context, f db.Handler) error {
	return m.transaction(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted}, f)
}
