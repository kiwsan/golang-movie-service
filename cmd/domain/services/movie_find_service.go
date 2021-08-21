package services

import (
	"fmt"

	"github.com/kiwsan/golang-movie-service/cmd/domain/exceptions"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieFindService interface {
	Find(id int64) (movie models.Movie, err error)
}

type MovieFindService struct {
	MovieRepository      repositories.IMovieRepository
	MovieRedisRepository repositories.IMovieRedisRepository
}

func (service *MovieFindService) Find(id int64) (movie models.Movie, err error) {

	movie, err = service.MovieRedisRepository.Get(id)
	logger.Info(fmt.Sprintf("Consulta Redis: %s", movie.Name))

	if err != nil {
		movie, err = service.MovieRepository.Find(id)
		if err != nil {
			err = exceptions.DataNotFound{ErrMessage: errorNotFoundRepository}
			logger.Error(errorRepository, err)
			return models.Movie{}, err
		}
		error := service.MovieRedisRepository.Set(movie)
		if error != nil {
			logger.Error(errorRepository, error)
		}
		return movie, nil
	}

	return movie, nil
}
