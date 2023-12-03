package urls

import (
	"context"
	"database/sql"
	"errors"
	"ozon_trainee_task/internal/common"

	"github.com/jmoiron/sqlx"
)

func NewRepository(dbConn *sqlx.DB) *repository {
	return &repository{dbConn: dbConn}
}

type repository struct {
	dbConn *sqlx.DB
}

const checkExistOriginalURL = `
	SELECT alias
	FROM url_mappings
	WHERE original_url = $1
`

const checkExistAliasQuery = `
	SELECT EXISTS(
		SELECT 1 
		FROM url_mappings 
		WHERE alias = $1
	);
`

const insertShortUrlQuery = `
	INSERT INTO url_mappings(original_url, alias) 
	VALUES ($1, $2)
`

func (r *repository) ShortenURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error) {
	var alias string

	tx, err := r.dbConn.BeginTxx(ctx, nil)

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

	alias = common.GenerateAlias()

	for {
		var aliasExist bool

		err = tx.GetContext(ctx, &aliasExist, checkExistAliasQuery, alias)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		if !aliasExist {
			break
		}

		alias = common.GenerateAlias()
	}

	_, err = tx.ExecContext(ctx, insertShortUrlQuery, in.OriginalURL, alias)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &ShortenURLOut{Alias: alias}, nil
}

func (r *repository) GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error) {
	//TODO implement me
	panic("implement me")
}
