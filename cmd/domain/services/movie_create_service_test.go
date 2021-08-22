package services_test

import (
	"errors"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/cmd/tests/builders"
	"github.com/kiwsan/golang-movie-service/cmd/tests/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	movieRepository      = new(mocks.MovieRepositoryMock)
	movieRedisRepository = new(mocks.MovieRedisRepositoryMock)
)

func TestWhenSendTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	movieRepository.On("Exist", movie.Name).Times(1).Return(false)
	movieRepository.On("Create", movie).Times(1).Return(nil)
	movieCreationService := services.MovieCreateService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.Add(movie)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedSendTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Exist", movie.Name).Times(1).Return(false)
	movieRepository.On("Create", movie).Times(1).Return(errorExpected)
	movieCreationService := services.MovieCreateService{
		MovieRepository: movieRepository,
	}

	err := movieCreationService.Add(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}

func TestWhenMovieExistThenShouldReturnError(t *testing.T) {
	movie := builders.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorExistRepository)
	movieRepository.On("Exist", movie.Name).Times(1).Return(true)
	movieCreationService := services.MovieCreateService{
		MovieRepository: movieRepository,
	}
	err := movieCreationService.Add(movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
