package queries

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieFindQuery interface {
	Handler(id int64) (movie models.Movie, err error)
}

type MovieFindQuery struct {
	MovieFindService services.IMovieFindService
}

func (query *MovieFindQuery) Handler(id int64) (movie models.Movie, err error) {

	movie, err = query.MovieFindService.Find(id)

	if err != nil {
		logger.Error(errorServiceMovieFind, err)
		return movie, err
	}
	return movie, err
}
