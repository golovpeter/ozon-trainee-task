package urls

import (
	"context"

	"github.com/golovpeter/ozon-trainee-task/internal/service/urls"
	"github.com/golovpeter/ozon-trainee-task/protos/gen/go/url_shortener"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *handler) GetOriginalURL(
	ctx context.Context,
	in *url_shortener.GetOriginalURLRequest,
) (*url_shortener.GetOriginalURLResponse, error) {
	out, err := s.urlService.GetOriginalURL(ctx, &urls.GetOriginalURLIn{ShortURL: in.ShortenedUrl})
	if err != nil {
		s.log.WithError(err).Error("failed to get original url")
		return nil, status.Errorf(codes.Internal, "failed to get original url")
	}

	return &url_shortener.GetOriginalURLResponse{OriginalUrl: out.OriginalURL}, nil
}
