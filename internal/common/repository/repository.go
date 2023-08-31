package repository

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urban-lib/logging/v2"
	"time"
)

type IRepository[T any] interface {
	FindOne(ctx context.Context, query string) (*T, error)
	Find(ctx context.Context, query string) ([]*T, error)
	FindWithDest(ctx context.Context, query string, dst any) error
	FindOneWithDest(ctx context.Context, query string, dst any) error
	Save(ctx context.Context, query string) (int, error)
	Close(ctx context.Context)
	Begin(ctx context.Context) (pgx.Tx, error)
	SaveWithTX(ctx context.Context, tx pgx.Tx, query string) (int, error)
}

type repository[T any] struct {
	connect *pgxpool.Pool
	timeout time.Duration
}

func (this *repository[T]) FindOne(ctx context.Context, query string) (*T, error) {
	object := new(T)
	err := SlowQueryRow(this.timeout, query, func() error {
		return pgxscan.Get(ctx, this.connect, object, query)
	})
	if err != nil {
		logging.Errorf("%s, query: %s", err.Error(), query)
		return nil, err
	}
	return object, nil
}

func (this *repository[T]) Find(ctx context.Context, query string) ([]*T, error) {
	var objects []*T
	err := SlowQueryRow(this.timeout, query, func() error {
		return pgxscan.Select(ctx, this.connect, &objects, query)
	})
	if err != nil {
		logging.Errorf("%s, query: %s", err.Error(), query)
		return nil, err
	}
	return objects, nil
}

func (this *repository[T]) FindWithDest(ctx context.Context, query string, dst any) error {
	err := SlowQueryRow(this.timeout, query, func() error {
		return pgxscan.Select(ctx, this.connect, dst, query)
	})
	if err != nil {
		logging.Errorf("%s, query: %s", err.Error(), query)
		return err
	}
	return nil
}

func (this *repository[T]) FindOneWithDest(ctx context.Context, query string, dst any) error {
	err := SlowQueryRow(this.timeout, query, func() error {
		return pgxscan.Get(ctx, this.connect, dst, query)
	})
	if err != nil {
		logging.Errorf("%s, query: %s", err.Error(), query)
		return err
	}
	return nil
}

func (this *repository[T]) Save(ctx context.Context, query string) (int, error) {
	var objectID int
	err := SlowQueryRow(this.timeout, query, func() error {
		return pgxscan.Select(ctx, this.connect, &objectID, query)
	})
	if err != nil {
		logging.Errorf("%s, query: %s", err.Error(), query)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return 0, ErrorRecordDuplicate
		}
		return 0, err
	}
	return objectID, nil
}

func (this *repository[T]) Close(ctx context.Context) {
	this.connect.Close()
}

func (this *repository[T]) Begin(ctx context.Context) (pgx.Tx, error) {
	return this.connect.Begin(ctx)
}

func (this *repository[T]) SaveWithTX(ctx context.Context, tx pgx.Tx, query string) (int, error) {
	var objectID int
	err := SlowQueryRow(this.timeout, query, func() error {
		return tx.QueryRow(ctx, query).Scan(&objectID)
	})
	if err != nil {
		logging.Errorf("%s, query: %s", err.Error(), query)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return 0, ErrorRecordDuplicate
		}
		return 0, err
	}
	return objectID, nil
}

func New[T any](connect *pgxpool.Pool, timeout time.Duration) IRepository[T] {
	return &repository[T]{connect: connect, timeout: timeout}
}
