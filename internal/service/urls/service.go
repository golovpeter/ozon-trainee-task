package urls

import (
	"context"

	"github.com/golovpeter/ozon-trainee-task/internal/common"
	"github.com/golovpeter/ozon-trainee-task/internal/repository/urls"
)

type service struct {
	urlsRepository urls.Repository
}

func NewService(urlsRepository urls.Repository) *service {
	return &service{urlsRepository: urlsRepository}
}

func (s *service) ShortenURL(ctx context.Context, in *ShortenUrlIn) (*ShortenURLOut, error) {
	newAlias := common.GenerateAlias(in.OriginalURL)

	url, err := s.urlsRepository.SaveShortenedURL(ctx, &urls.ShortenUrlIn{
		OriginalURL: in.OriginalURL,
		Alias:       newAlias,
	})

	if err != nil {
		return nil, err
	}

	return &ShortenURLOut{Alias: url.Alias}, nil
}

func (s *service) GetOriginalURL(ctx context.Context, in *GetOriginalURLIn) (*GetOriginalURlOut, error) {
	alias := common.GetAlias(in.ShortURL)

	out, err := s.urlsRepository.GetOriginalURL(ctx, &urls.GetOriginalURLIn{ShortURL: alias})
	if err != nil {
		return nil, err
	}

	return &GetOriginalURlOut{OriginalURL: out.OriginalURL}, nil
}
