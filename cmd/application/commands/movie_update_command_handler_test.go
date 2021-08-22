package commands_test

import (
	"errors"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/application/commands"
	"github.com/kiwsan/golang-movie-service/cmd/tests/builders"
	"github.com/kiwsan/golang-movie-service/cmd/tests/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	movieUpdateServiceMock = new(mocks.MovieUpdateServiceMock)
)

func TestWhenAllBeOKUpdatedMovieThenReturnNilError(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	movieUpdateServiceMock.On("Update", movie.Id, movie).Return(nil).Once()
	movieUpdate := commands.MovieUpdateCommand{
		MovieUpdateService: movieUpdateServiceMock,
	}

	err := movieUpdate.Handler(movie.Id, movie)

	assert.Nil(t, err)
	movieUpdateServiceMock.AssertExpectations(t)
}
func TestWhenFailedUpdatedParkingThenReturnError(t *testing.T) {
	movie := builders.NewMovieDataBuilder().Build()
	expectedErrorMessage := errors.New("error getting repository information")
	movieUpdateServiceMock.On("Update", movie.Id, movie).Return(expectedErrorMessage).Once()
	movieUpdate := commands.MovieUpdateCommand{
		MovieUpdateService: movieUpdateServiceMock,
	}

	err := movieUpdate.Handler(movie.Id, movie)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieUpdateServiceMock.AssertExpectations(t)
}
