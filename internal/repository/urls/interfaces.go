package urls

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=urls Urls
type Repository interface {
	SaveShortenedURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error)
	GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error)
}

type Database interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}
