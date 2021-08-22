package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieCreationMock struct {
	mock.Mock
}

func (mock *MovieCreationMock) Handler(movie models.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}
