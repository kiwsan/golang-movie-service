package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieRedisRepositoryMock struct {
	mock.Mock
}

func (mock *MovieRedisRepositoryMock) Set(movie models.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}

func (mock *MovieRedisRepositoryMock) Get(id int64) (movie models.Movie, err error) {
	args := mock.Called(id)
	return args.Get(0).(models.Movie), args.Error(1)
}
