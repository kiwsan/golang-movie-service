package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieFindAllMock struct {
	mock.Mock
}

func (mock *MovieFindAllMock) Handler() (movies []models.Movie, err error) {
	args := mock.Called()
	return args.Get(0).([]models.Movie), args.Error(1)
}
