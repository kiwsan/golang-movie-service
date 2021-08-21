package queries

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieFindAllQuery interface {
	Handler() (movieLots []models.Movie, err error)
}

type MovieFindAllQuery struct {
	MovieRepository repositories.IMovieRepository
}

func (query *MovieFindAllQuery) Handler() (movieLots []models.Movie, err error) {

	movieLots, err = query.MovieRepository.FindAll()

	if err != nil {
		logger.Error(errorServiceMovieList, err)
		return movieLots, err
	}
	return movieLots, err
}
