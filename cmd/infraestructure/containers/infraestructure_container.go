package containers

import (
	"github.com/kiwsan/golang-movie-service/cmd/api/controllers"
	irepository "github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/configs"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/repositories"
)

func GetHealthCheckController() *controllers.HealthCheckController {
	return &controllers.HealthCheckController{}
}

func GetMovieCreateController() *controllers.MovieCreateController {
	return &controllers.MovieCreateController{IMovieCreateCommand: GetMovieCreationAccessApplication()}
}

func GetMovieUpdateController() *controllers.MovieUpdateController {
	return &controllers.MovieUpdateController{IMovieUpdateCommand: GetMovieUpdateAccessApplication()}
}

func GetMovieDeleteController() *controllers.MovieDeleteController {
	return &controllers.MovieDeleteController{IMovieDeleteCommand: GetMovieDeleteAccessApplication()}
}

func GetMovieFindAllController() *controllers.MovieFindAllController {
	return &controllers.MovieFindAllController{IMovieFindAllQuery: GetMovieFindAllAccessApplication()}
}

func GetMovieFindController() *controllers.MovieFindController {
	return &controllers.MovieFindController{IMovieFindQuery: GetMovieFindAccessApplication()}
}

func getMovieRepository() irepository.IMovieRepository {
	return &repositories.MovieMySqlRepository{
		Db: configs.GetDatabaseInstance(),
	}
}

func getMovieRedisRepository() irepository.IMovieRedisRepository {
	return &repositories.MovieRedisRepository{
		Db: configs.GetClient(),
	}
}
