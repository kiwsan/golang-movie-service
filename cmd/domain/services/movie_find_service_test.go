package services_test

import (
	"errors"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/cmd/tests/builders"
	"github.com/stretchr/testify/assert"
)

func TestWhenGetTheMovieFindFromRepositoryThenShouldReturnOk(t *testing.T) {
	movieB := builders.NewMovieDataBuilder().Build()
	movieR := models.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieRedisRepository.On("Get", movieB.Id).Times(1).Return(movieR, errorExpected)
	movieRepository.On("Find", movieB.Id).Times(1).Return(movieB, nil)
	movieRedisRepository.On("Set", movieB).Times(1).Return(nil)
	movieFindService := services.MovieFindService{
		MovieRepository:      movieRepository,
		MovieRedisRepository: movieRedisRepository,
	}

	movie, err := movieFindService.Find(movieB.Id)

	assert.Nil(t, err)
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}
func TestWhenFailedGetTheMovieFromRepositoryThenShouldReturnError(t *testing.T) {

	movieB := models.Movie{}
	errorExpected := errors.New(errorFindRepository)
	movieR := models.Movie{}
	movieRedisRepository.On("Get", movieB.Id).Times(1).Return(movieR, errorExpected)
	movieRepository.On("Find", movieB.Id).Times(1).Return(movieB, errorExpected)
	movieFindService := services.MovieFindService{
		MovieRepository:      movieRepository,
		MovieRedisRepository: movieRedisRepository,
	}

	movie, err := movieFindService.Find(0)

	assert.NotNil(t, err)
	assert.EqualError(t, errorExpected, err.Error())
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}

func TestWhenGetTheMovieFindFromRedisRepositoryThenShouldReturnOk(t *testing.T) {
	movieB := builders.NewMovieDataBuilder().Build()
	movieRedisRepository.On("Get", movieB.Id).Times(1).Return(movieB, nil)
	movieFindService := services.MovieFindService{
		MovieRepository:      movieRepository,
		MovieRedisRepository: movieRedisRepository,
	}

	movie, err := movieFindService.Find(movieB.Id)

	assert.Nil(t, err)
	assert.Equal(t, movieB, movie)
	movieRepository.AssertExpectations(t)
}
