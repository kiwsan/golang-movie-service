package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieUpdateMock struct {
	mock.Mock
}

func (mock *MovieUpdateMock) Handler(id int64, movie models.Movie) (err error) {
	args := mock.Called(id, movie)
	return args.Error(0)
}
