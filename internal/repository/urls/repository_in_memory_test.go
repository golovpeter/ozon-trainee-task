package urls

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InMemorySuite struct {
	suite.Suite

	repository *repositoryInMemory
}

func (s *InMemorySuite) SetupTest() {
	s.repository = NewRepositoryInMemory()

	s.repository.data["abc"] = "https://example.com"
}

func TestRunInMemorySuite(t *testing.T) {
	suite.Run(t, new(InMemorySuite))
}

func (s *InMemorySuite) TestSaveShortenedURLSaveNewURL() {
	input := &ShortenUrlIn{Alias: "alias", OriginalURL: "https://example.com"}
	expectedOutput := &ShortenURLOut{Alias: "alias"}

	output, err := s.repository.SaveShortenedURL(context.TODO(), input)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedOutput.Alias, output.Alias)
}

func (s *InMemorySuite) TestGetOriginalURLSuccess() {
	input := &GetOriginalURLIn{ShortURL: "abc"}
	expectedOutput := &GetOriginalURlOut{"https://example.com"}

	output, err := s.repository.GetOriginalURL(context.TODO(), input)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedOutput.OriginalURL, output.OriginalURL)
}

func (s *InMemorySuite) TestGetOriginalURLAliasNotFound() {
	input := &GetOriginalURLIn{ShortURL: "unknown"}

	_, err := s.repository.GetOriginalURL(context.TODO(), input)

	assert.Error(s.T(), err)
}
