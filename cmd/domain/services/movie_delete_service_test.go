package services_test

import (
	"errors"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/cmd/tests/builders"
	"github.com/stretchr/testify/assert"
)

func TestWhenDeleteTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Delete", movie.Id).Times(1).Return(nil)
	movieDeleteService := services.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.Remove(movie.Id)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}

func TestWhenFailedSendToDeleteTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Delete", movie.Id).Times(1).Return(errorExpected)
	movieDeleteService := services.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.Remove(movie.Id)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}

func TestWhenFailedFindTheMovieToDeleteThenShouldReturnError(t *testing.T) {

	movie := models.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, errorExpected)
	movieDeleteService := services.MovieDeleteService{
		MovieRepository: movieRepository,
	}
	err := movieDeleteService.Remove(movie.Id)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
