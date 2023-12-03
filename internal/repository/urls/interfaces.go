package urls

import "context"

type Repository interface {
	ShortenURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error)
	GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error)
}
