package services

import (
	"errors"

	"github.com/kiwsan/golang-movie-service/cmd/domain/exception"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieCreateService interface {
	Add(movie models.Movie) (err error)
}

type MovieCreateService struct {
	MovieRepository repositories.MovieRepository
}

func (service *MovieCreateService) Add(movie models.Movie) (err error) {

	exist := service.MovieRepository.Exist(movie.Name)
	if exist {
		exp := exception.DataDuplicity{ErrMessage: errorExistRepository}
		logger.Error(errorExistRepository, exp)
		return exp
	}

	err = service.MovieRepository.Create(movie)
	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
