package mocks

import "github.com/stretchr/testify/mock"

type MovieDeleteMock struct {
	mock.Mock
}

func (mock *MovieDeleteMock) Handler(id int64) (err error) {
	args := mock.Called(id)
	return args.Error(0)
}
