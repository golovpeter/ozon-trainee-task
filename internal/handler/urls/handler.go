package urls

import (
	urlsService "github.com/golovpeter/ozon-trainee-task/internal/service/urls"
	"github.com/golovpeter/ozon-trainee-task/protos/gen/go/url_shortener"
	"github.com/sirupsen/logrus"
)

type handler struct {
	url_shortener.UnimplementedUrlShortenerServer

	log        *logrus.Logger
	urlService urlsService.UrlService
}

func NewHandler(log *logrus.Logger, urlService urlsService.UrlService) *handler {
	return &handler{
		log:        log,
		urlService: urlService,
	}
}
