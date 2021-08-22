package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieFindAllServiceMock struct {
	mock.Mock
}

func (mock *MovieFindAllServiceMock) FindAll() (movies []models.Movie, err error) {
	args := mock.Called()
	return args.Get(0).([]models.Movie), args.Error(1)
}
