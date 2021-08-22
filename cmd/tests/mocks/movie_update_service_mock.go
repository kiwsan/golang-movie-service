package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieUpdateServiceMock struct {
	mock.Mock
}

func (mock *MovieUpdateServiceMock) Update(id int64, movie models.Movie) error {
	args := mock.Called(id, movie)
	return args.Error(0)
}
