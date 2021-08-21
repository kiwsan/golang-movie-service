package containers

import "github.com/kiwsan/golang-movie-service/cmd/domain/services"

func getCreateMovieService() services.IMovieCreateService {
	return &services.MovieCreateService{
		MovieRepository: getMovieRepository(),
	}
}

func getUpdateMovieService() services.IMovieUpdateService {
	return &services.MovieUpdateService{
		MovieRepository: getMovieRepository(),
	}
}

func getDeleteMovieService() services.IMovieDeleteService {
	return &services.MovieDeleteService{
		MovieRepository: getMovieRepository(),
	}
}

func getFindMovieService() services.IMovieFindService {
	return &services.MovieFindService{
		MovieRepository:      getMovieRepository(),
		MovieRedisRepository: getMovieRedisRepository(),
	}
}
