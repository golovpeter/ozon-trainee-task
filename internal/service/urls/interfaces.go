package urls

import "context"

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=urls Urls
type UrlService interface {
	ShortenURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error)
	GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error)
}
