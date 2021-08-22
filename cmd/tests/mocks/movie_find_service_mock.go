package mocks

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/stretchr/testify/mock"
)

type MovieFindServiceMock struct {
	mock.Mock
}

func (mock *MovieFindServiceMock) Find(id int64) (movie models.Movie, err error) {
	args := mock.Called(id)
	return args.Get(0).(models.Movie), args.Error(1)
}
