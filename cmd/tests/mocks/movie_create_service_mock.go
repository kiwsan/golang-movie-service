package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieCreationServiceMock struct {
	mock.Mock
}

func (mock *MovieCreationServiceMock) Add(movie models.Movie) error {
	args := mock.Called(movie)
	return args.Error(0)
}
