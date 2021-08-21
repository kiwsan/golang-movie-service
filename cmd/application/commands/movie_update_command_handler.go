package commands

import (
	"fmt"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieUpdateCommand interface {
	Handler(id int64, movie models.Movie) (err error)
}

type MovieUpdateCommand struct {
	MovieUpdateService services.IMovieUpdateService
}

func (command *MovieUpdateCommand) Handler(id int64, movie models.Movie) (err error) {

	err = command.MovieUpdateService.Update(id, movie)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceMovie, movie.Name), err)
		return err
	}
	return err
}
