package services_test

import (
	"errors"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/cmd/tests/builders"
	"github.com/stretchr/testify/assert"
)

func TestWhenSendToUpdateTheMovieToRepositoryThenShouldReturnOk(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Update", movie.Id, movie).Times(1).Return(nil)
	movieUpdateService := services.MovieUpdateService{
		MovieRepository: movieRepository,
	}
	err := movieUpdateService.Update(movie.Id, movie)

	assert.Nil(t, err)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedSendToUpdateTheMovieToRepositoryThenShouldReturnError(t *testing.T) {

	movie := builders.NewMovieDataBuilder().Build()
	errorExpected := errors.New(errorRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, nil)
	movieRepository.On("Update", movie.Id, movie).Times(1).Return(errorExpected)
	movieUpdateService := services.MovieUpdateService{
		MovieRepository: movieRepository,
	}
	err := movieUpdateService.Update(movie.Id, movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}

func TestWhenFailedFindTheMovieToUpdateThenShouldReturnError(t *testing.T) {

	movie := models.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieRepository.On("Find", movie.Id).Times(1).Return(movie, errorExpected)
	movieUpdateService := services.MovieUpdateService{
		MovieRepository: movieRepository,
	}
	err := movieUpdateService.Update(movie.Id, movie)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	movieRepository.AssertExpectations(t)
}
