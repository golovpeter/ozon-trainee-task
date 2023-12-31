package urls

import (
	"context"
	"database/sql"
	"errors"
)

type repositoryPostgres struct {
	db Database
}

func NewRepositoryPostgres(db Database) *repositoryPostgres {
	return &repositoryPostgres{db: db}
}

const checkExistOriginalURL = `
	SELECT alias
	FROM url_mappings
	WHERE original_url = $1
`

const insertShortUrlQuery = `
	INSERT INTO url_mappings(original_url, alias) 
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING 
`

func (r *repositoryPostgres) SaveShortenedURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error) {
	var alias string

	tx, err := r.db.BeginTxx(ctx, nil)

	defer func() {
		if rec := recover(); rec != nil {
			_ = tx.Rollback()
		}
	}()

	err = tx.GetContext(ctx, &alias, checkExistOriginalURL, in.OriginalURL)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		_ = tx.Rollback()
		return nil, err
	}

	if alias != "" {
		return &ShortenURLOut{Alias: alias}, nil
	}

	_, err = tx.ExecContext(ctx, insertShortUrlQuery, in.OriginalURL, in.Alias)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &ShortenURLOut{Alias: in.Alias}, nil
}

const getOriginURLQuery = `
	SELECT original_url
	FROM url_mappings
	WHERE alias = $1
`

func (r *repositoryPostgres) GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error) {
	var originalURL string

	err := r.db.GetContext(ctx, &originalURL, getOriginURLQuery, in.ShortURL)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("original url not found")
	}

	if err != nil {
		return nil, err
	}

	return &GetOriginalURlOut{OriginalURL: originalURL}, nil
}
