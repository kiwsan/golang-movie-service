package services

import (
	"errors"

	"github.com/kiwsan/golang-movie-service/cmd/domain/exceptions"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieUpdateService interface {
	Update(id int64, movie models.Movie) (err error)
}

type MovieUpdateService struct {
	MovieRepository repositories.IMovieRepository
}

func (service *MovieUpdateService) Update(id int64, movie models.Movie) (err error) {

	_, exist := service.MovieRepository.Find(id)
	if exist != nil {
		err = exceptions.DataNotFound{ErrMessage: errorNotFoundRepository}
		logger.Error(errorRepository, err)
		return err
	}

	err = service.MovieRepository.Update(id, movie)

	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
