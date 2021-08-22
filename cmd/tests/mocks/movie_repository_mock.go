package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieRepositoryMock struct {
	mock.Mock
}

func (mock *MovieRepositoryMock) Create(movie models.Movie) (err error) {
	args := mock.Called(movie)
	return args.Error(0)
}

func (mock *MovieRepositoryMock) Exist(name string) (exist bool) {
	args := mock.Called(name)
	return args.Get(0).(bool)
}

func (mock *MovieRepositoryMock) Find(id int64) (movie models.Movie, err error) {
	args := mock.Called(id)
	return args.Get(0).(models.Movie), args.Error(1)
}

func (mock *MovieRepositoryMock) FindAll() (movies []models.Movie, err error) {
	args := mock.Called()
	return args.Get(0).([]models.Movie), args.Error(1)
}

func (mock *MovieRepositoryMock) Update(id int64, movie models.Movie) (err error) {
	args := mock.Called(id, movie)
	return args.Error(0)
}

func (mock *MovieRepositoryMock) Delete(id int64) (err error) {
	args := mock.Called(id)
	return args.Error(0)
}
