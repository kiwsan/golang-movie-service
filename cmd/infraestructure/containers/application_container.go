package containers

import (
	"github.com/kiwsan/golang-movie-service/cmd/application/commands"
	"github.com/kiwsan/golang-movie-service/cmd/application/queries"
)

func GetMovieCreationAccessApplication() commands.IMovieCreateCommand {
	return &commands.MovieCreateCommand{MovieCreateService: getCreateMovieService()}
}

func GetMovieUpdateAccessApplication() commands.IMovieUpdateCommand {
	return &commands.MovieUpdateCommand{MovieUpdateService: getUpdateMovieService()}
}

func GetMovieDeleteAccessApplication() commands.IMovieDeleteCommand {
	return &commands.MovieDeleteCommand{MovieDeleteService: getDeleteMovieService()}
}

func GetMovieFindAllAccessApplication() queries.IMovieFindAllQuery {
	return &queries.MovieFindAllQuery{MovieRepository: getMovieRepository()}
}

func GetMovieFindAccessApplication() queries.IMovieFindQuery {
	return &queries.MovieFindQuery{MovieFindService: getFindMovieService()}
}
