package urls

import (
	"context"
	"errors"
	"testing"

	"github.com/golovpeter/ozon-trainee-task/internal/service/urls"
	"github.com/golovpeter/ozon-trainee-task/protos/gen/go/url_shortener"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockUrlService *urls.MockUrlService

	handler *handler
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockUrlService = urls.NewMockUrlService(ts.ctrl)

	ts.handler = NewHandler(logrus.New(), ts.mockUrlService)
}

func (ts *TestSuite) TestShortenURLSuccess() {
	testShortenedURL := "https://example.com/alias"
	testAlias := "alias"

	input := &url_shortener.ShortenURLRequest{
		OriginalUrl: "https://example.com",
	}

	customMetadata := metadata.Pairs(
		":authority", "https://example.com",
	)
	ctx := metadata.NewIncomingContext(context.Background(), customMetadata)

	ts.mockUrlService.EXPECT().
		ShortenURL(ctx, &urls.ShortenUrlIn{
			OriginalURL: input.OriginalUrl,
		}).Return(&urls.ShortenURLOut{
		Alias: testAlias,
	}, nil).
		Times(1)

	output, err := ts.handler.ShortenURL(ctx, input)

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), output.ShortenedUrl, testShortenedURL)
}

func (ts *TestSuite) TestShortenURLServiceError() {
	input := &url_shortener.ShortenURLRequest{
		OriginalUrl: "https://example.com",
	}

	customMetadata := metadata.Pairs(
		":authority", "https://example.com",
	)
	ctx := metadata.NewIncomingContext(context.Background(), customMetadata)

	ts.mockUrlService.EXPECT().
		ShortenURL(ctx, gomock.Any()).
		Return(nil, errors.New("service error")).
		Times(1)

	_, err := ts.handler.ShortenURL(ctx, input)

	assert.Error(ts.T(), err)
}

func (ts *TestSuite) TestShortenURLEmptyMetadata() {
	input := &url_shortener.ShortenURLRequest{
		OriginalUrl: "https://example.com",
	}

	ts.mockUrlService.EXPECT().
		ShortenURL(context.TODO(), gomock.Any()).
		Return(nil, nil).
		Times(1)

	_, err := ts.handler.ShortenURL(context.TODO(), input)

	assert.Error(ts.T(), err)
}

func (ts *TestSuite) TestGetOriginalURLSuccess() {
	testOriginalURL := "https://example.com"

	input := &url_shortener.GetOriginalURLRequest{
		ShortenedUrl: "https://example.com/alias",
	}

	expectedOutput := &urls.GetOriginalURlOut{
		OriginalURL: testOriginalURL,
	}

	ts.mockUrlService.EXPECT().
		GetOriginalURL(context.TODO(), &urls.GetOriginalURLIn{
			ShortURL: input.ShortenedUrl,
		}).Return(expectedOutput, nil)

	output, err := ts.handler.GetOriginalURL(context.TODO(), input)

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), testOriginalURL, output.OriginalUrl)
}

func (ts *TestSuite) TestGetOriginalURLServiceError() {
	input := &url_shortener.GetOriginalURLRequest{
		ShortenedUrl: "https://example.com/alias",
	}

	ts.mockUrlService.EXPECT().
		GetOriginalURL(context.TODO(), gomock.Any()).
		Return(nil, errors.New("service error")).
		Times(1)

	_, err := ts.handler.GetOriginalURL(context.TODO(), input)

	assert.Error(ts.T(), err)
}
