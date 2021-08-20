package services

import (
	"errors"

	"github.com/kiwsan/golang-movie-service/cmd/domain/exception"
	"github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieDeleteService interface {
	Remove(id int64) (err error)
}

type MovieDeleteService struct {
	MovieRepository repositories.MovieRepository
}

func (service *MovieDeleteService) Remove(id int64) (err error) {

	_, exist := service.MovieRepository.Find(id)
	if exist != nil {
		err = exception.DataNotFound{ErrMessage: errorNotFoundRepository}
		return err
	}

	err = service.MovieRepository.Delete(id)
	if err != nil {
		err = errors.New(errorRepository)
		logger.Error(errorRepository, err)
		return err
	}

	return err
}
