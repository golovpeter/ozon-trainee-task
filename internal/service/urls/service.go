package urls

import (
	"context"
	"ozon_trainee_task/internal/repository/urls"
)

type service struct {
	urlsRepository urls.Repository
}

func NewService(urlsRepository urls.Repository) *service {
	return &service{urlsRepository: urlsRepository}
}

func (s *service) ShortenURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error) {
	url, err := s.urlsRepository.ShortenURL(ctx, &urls.ShortenUrlIn{
		OriginalURL: in.OriginalURL,
	})

	if err != nil {
		return nil, err
	}

	return &ShortenURLOut{Alias: url.Alias}, nil
}

func (s *service) GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error) {
	//TODO implement me
	panic("implement me")
}
