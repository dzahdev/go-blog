package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type key string

const (
	TxKey key = "tx"
)

// Handler is a function that can be executed in a transaction.
type Handler func(ctx context.Context) error

// Client is a database client.
type Client interface {
	DB() DB
	Close() error
}

// TxManager is a transaction manager.
type TxManager interface {
	ReadCommited(ctx context.Context, f Handler) error
}

// Query is a query.
type Query struct {
	Name     string
	QueryRow string
}

type Transactor interface {
	Begin(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, res interface{}, query Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, res interface{}, query Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
