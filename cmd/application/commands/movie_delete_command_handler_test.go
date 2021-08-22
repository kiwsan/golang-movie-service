package commands_test

import (
	"errors"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/application/commands"
	"github.com/kiwsan/golang-movie-service/cmd/tests/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	movieDeleteServiceMock = new(mocks.MovieDeleteServiceMock)
)

func TestWhenAllBeOKDeletedMovieThenReturnNilError(t *testing.T) {
	const id int64 = 5
	movieDeleteServiceMock.On("Delete", id).Return(nil).Once()
	movieDelete := commands.MovieDeleteCommand{
		MovieDeleteService: movieDeleteServiceMock,
	}

	err := movieDelete.Handler(id)

	assert.Nil(t, err)
	movieDeleteServiceMock.AssertExpectations(t)
}
func TestWhenFailedDeletedParkingThenReturnError(t *testing.T) {
	const id int64 = 5
	expectedErrorMessage := errors.New("error getting repository information")
	movieDeleteServiceMock.On("Delete", id).Return(expectedErrorMessage).Once()
	movieDelete := commands.MovieDeleteCommand{
		MovieDeleteService: movieDeleteServiceMock,
	}

	err := movieDelete.Handler(id)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieDeleteServiceMock.AssertExpectations(t)
}
