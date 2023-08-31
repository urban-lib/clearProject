package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/urban-lib/logging/v2"
	"time"
)

func SlowQueryExec(limit time.Duration, query string, caller func() (pgconn.CommandTag, error)) (pgconn.CommandTag, error) {
	start := time.Now()
	defer func() {
		if time.Since(start).Milliseconds() >= limit.Milliseconds() {
			logging.Warn("SLOW QUERY ", time.Since(start).Milliseconds(), "ms. ", query)
		}
	}()
	return caller()
}

func SlowQuery(limit time.Duration, query string, caller func() (pgx.Rows, error)) (pgx.Rows, error) {
	start := time.Now()
	defer func() {
		if time.Since(start).Milliseconds() >= limit.Milliseconds() {
			logging.Warn("SLOW QUERY ", time.Since(start).Milliseconds(), "ms. ", query)
		}
	}()
	return caller()
}

func SlowQueryRow(limit time.Duration, query string, caller func() error) error {
	start := time.Now()
	defer func() {
		if time.Since(start).Milliseconds() >= limit.Milliseconds() {
			logging.Warn("SLOW QUERY ", time.Since(start).Milliseconds(), "ms. ", query)
		}
	}()
	return caller()
}
