package urls

import (
	"context"
	"errors"
	"testing"

	"github.com/golovpeter/ozon-trainee-task/internal/common"
	"github.com/golovpeter/ozon-trainee-task/internal/repository/urls"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockUrlsRepository *urls.MockRepository

	service UrlService
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	ts.ctrl = gomock.NewController(ts.T())

	ts.mockUrlsRepository = urls.NewMockRepository(ts.ctrl)

	ts.service = NewService(
		ts.mockUrlsRepository,
	)
}

func (ts *TestSuite) TestShortenURLSuccess() {
	testAlias := "alias"

	generateAlias = func(input string) string {
		return testAlias
	}

	defer func() {
		generateAlias = common.GenerateAlias
	}()

	input := &ShortenUrlIn{
		OriginalURL: "https://example.com",
	}

	expectedOutput := &urls.ShortenURLOut{Alias: testAlias}

	ts.mockUrlsRepository.EXPECT().
		SaveShortenedURL(context.TODO(), &urls.ShortenUrlIn{
			OriginalURL: input.OriginalURL,
			Alias:       testAlias,
		}).
		Times(1).
		Return(expectedOutput, nil)

	output, err := ts.service.ShortenURL(context.TODO(), input)

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), expectedOutput.Alias, output.Alias)
}

func (ts *TestSuite) TestShortenURLServiceError() {
	input := &ShortenUrlIn{
		OriginalURL: "https://example.com",
	}

	ts.mockUrlsRepository.EXPECT().
		SaveShortenedURL(context.TODO(), gomock.Any()).
		Times(1).
		Return(nil, errors.New("service error"))

	_, err := ts.service.ShortenURL(context.TODO(), input)

	assert.Error(ts.T(), err)
}

func (ts *TestSuite) TestGetOriginalURLSuccess() {
	testAlias := "alias"

	input := &GetOriginalURLIn{
		ShortURL: "https://exmaple.com/alias",
	}

	expectedOutput := &urls.GetOriginalURlOut{OriginalURL: "google.com"}

	ts.mockUrlsRepository.EXPECT().
		GetOriginalURL(context.TODO(), &urls.GetOriginalURLIn{
			ShortURL: testAlias,
		}).
		Times(1).
		Return(expectedOutput, nil)

	output, err := ts.service.GetOriginalURL(context.TODO(), input)

	assert.NoError(ts.T(), err)
	assert.Equal(ts.T(), expectedOutput.OriginalURL, output.OriginalURL)
}

func (ts *TestSuite) TestGetOriginalURlError() {
	input := &GetOriginalURLIn{
		ShortURL: "https://exmaple.com/alias",
	}

	ts.mockUrlsRepository.EXPECT().
		GetOriginalURL(context.TODO(), gomock.Any()).
		Times(1).
		Return(nil, errors.New("service error"))

	_, err := ts.service.GetOriginalURL(context.TODO(), input)

	assert.Error(ts.T(), err)
}
