package containers

import (
	"github.com/go-playground/validator"
	"github.com/kiwsan/golang-movie-service/cmd/api/controllers"
	irepository "github.com/kiwsan/golang-movie-service/cmd/domain/repositories"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/configs"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/repositories"
)

func GetHealthCheckController() *controllers.HealthCheckController {
	return &controllers.HealthCheckController{}
}

func GetMovieCreateController(validate *validator.Validate) *controllers.MovieCreateController {
	return &controllers.MovieCreateController{
		InputValidate:       validate,
		IMovieCreateCommand: GetMovieCreationAccessApplication(),
	}
}

func GetMovieUpdateController(validate *validator.Validate) *controllers.MovieUpdateController {
	return &controllers.MovieUpdateController{
		InputValidate:       validate,
		IMovieUpdateCommand: GetMovieUpdateAccessApplication(),
	}
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
