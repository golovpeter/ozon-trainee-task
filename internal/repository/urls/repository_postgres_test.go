package urls

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type PostgresSuite struct {
	suite.Suite

	ctl *gomock.Controller

	dbMock *MockDatabase

	repository *repositoryPostgres
}

func (s *PostgresSuite) SetupTest() {
	s.ctl = gomock.NewController(s.T())

	s.dbMock = NewMockDatabase(s.ctl)
	s.repository = NewRepositoryPostgres(s.dbMock)
}

func TestRunPostgresSuite(t *testing.T) {
	suite.Run(t, new(PostgresSuite))
}

func (s *PostgresSuite) TestGetOriginalURLSuccess() {
	var originalURL string
	input := &GetOriginalURLIn{ShortURL: "abc"}

	s.dbMock.EXPECT().
		GetContext(context.TODO(), gomock.Any(), getOriginURLQuery, input.ShortURL).
		Times(1).
		Return(nil)

	output, err := s.repository.GetOriginalURL(context.TODO(), input)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), originalURL, output.OriginalURL)
}

func (s *PostgresSuite) TestGetOriginalURLOriginalURlNotFound() {
	input := &GetOriginalURLIn{ShortURL: "abc"}

	s.dbMock.EXPECT().
		GetContext(context.TODO(), gomock.Any(), getOriginURLQuery, input.ShortURL).
		Times(1).
		Return(sql.ErrNoRows)

	_, err := s.repository.GetOriginalURL(context.TODO(), input)

	assert.Error(s.T(), err)
}

func (s *PostgresSuite) TestGetOriginalURLGetContextError() {
	input := &GetOriginalURLIn{ShortURL: "abc"}

	s.dbMock.EXPECT().
		GetContext(context.TODO(), gomock.Any(), getOriginURLQuery, input.ShortURL).
		Times(1).
		Return(errors.New("another error"))

	_, err := s.repository.GetOriginalURL(context.TODO(), input)

	assert.Error(s.T(), err)
}
