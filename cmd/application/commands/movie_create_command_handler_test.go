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
	movieCreationServiceMock = new(mocks.MovieCreationServiceMock)
)

func TestWhenAllBeOKCreatingMovieThenReturnNilError(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	movieCreationServiceMock.On("Create", movie).Return(nil).Once()
	movieCreation := commands.MovieCreateCommand{
		MovieCreateService: movieCreationServiceMock,
	}

	err := movieCreation.Handler(movie)

	assert.Nil(t, err)
	movieCreationServiceMock.AssertExpectations(t)
}

func TestWhenFailedCreatingParkingThenReturnError(t *testing.T) {
	movie := builders.NewMovieDataBuilder().Build()
	expectedErrorMessage := errors.New("error getting repository information")
	movieCreationServiceMock.On("Create", movie).Return(expectedErrorMessage).Once()
	movieCreation := commands.MovieCreateCommand{
		MovieCreateService: movieCreationServiceMock,
	}

	err := movieCreation.Handler(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErrorMessage.Error())
	movieCreationServiceMock.AssertExpectations(t)
}
