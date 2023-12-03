package urls

import (
	"context"
	"fmt"

	"github.com/golovpeter/ozon-trainee-task/internal/service/urls"
	"github.com/golovpeter/ozon-trainee-task/protos/gen/go/url_shortener"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *handler) ShortenURL(
	ctx context.Context,
	in *url_shortener.ShortenURLRequest,
) (*url_shortener.ShortenURLResponse, error) {
	out, err := s.urlService.ShortenURL(ctx, &urls.ShortenUrlIn{OriginalURL: in.GetOriginalUrl()})

	if err != nil {
		s.log.WithError(err).Error("failed to shorten url")
		return nil, status.Error(codes.Internal, "failed to short url")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.log.Error("not found metadata")
		return nil, status.Error(codes.Unavailable, "not found metadata")
	}

	host := md.Get(":authority")[0]

	return &url_shortener.ShortenURLResponse{
		ShortenedUrl: fmt.Sprintf("%s/%s", host, out.Alias),
	}, nil
}
