package commands

import (
	"fmt"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieCreateCommand interface {
	Handler(movie models.Movie) (err error)
}

type MovieCreateCommand struct {
	MovieCreateService services.IMovieCreateService
}

func (command *MovieCreateCommand) Handler(movie models.Movie) (err error) {

	err = command.MovieCreateService.Add(movie)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceMovie, movie.Name), err)
		return err
	}
	return err
}
