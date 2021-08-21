package commands

import (
	"fmt"

	"github.com/kiwsan/golang-movie-service/cmd/domain/services"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type IMovieDeleteCommand interface {
	Handler(id int64) (err error)
}

type MovieDeleteCommand struct {
	MovieDeleteService services.IMovieDeleteService
}

func (command *MovieDeleteCommand) Handler(id int64) (err error) {

	err = command.MovieDeleteService.Remove(id)
	if err != nil {
		logger.Error(fmt.Sprintf(errorServiceMovie, id), err)
		return err
	}
	return err
}
