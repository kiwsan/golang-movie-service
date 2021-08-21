package mappers

import (
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/entities"
)

func MovieToMovieEntity(movie models.Movie) entities.MovieEntity {
	var entity entities.MovieEntity

	entity.Id = movie.Id
	entity.Director = movie.Director
	entity.Name = movie.Name
	entity.Stars = movie.Stars
	entity.Writer = movie.Writer

	return entity
}

func MovieEntityToMovie(entity entities.MovieEntity) models.Movie {
	var model models.Movie

	model.Id = entity.Id
	model.Director = entity.Director
	model.Name = entity.Name
	model.Stars = entity.Stars
	model.Writer = entity.Writer

	return model
}

func MoviesEntitiesToMovies(entities []entities.MovieEntity) []models.Movie {
	var movies []models.Movie

	for _, entity := range entities {
		movie := MovieEntityToMovie(entity)
		movies = append(movies, movie)
	}

	return movies
}
