package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/entities"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/mappers"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

type MovieMySqlRepository struct {
	Db *gorm.DB
}

func (repo *MovieMySqlRepository) Create(movie models.Movie) error {
	var entity = mappers.MovieToMovieEntity(movie)
	if err := repo.Db.Create(&entity).Error; err != nil {
		logger.Error(fmt.Sprintf("can't work with %s", entity.Name), err)
		return fmt.Errorf(fmt.Sprintf("can't work with %s", entity.Name))
	}
	movie.Id = entity.Id
	return nil
}

func (repo *MovieMySqlRepository) Exist(name string) bool {
	var entity entities.MovieEntity

	return !(repo.Db.Where(&entities.MovieEntity{Name: name}).Find(&entity).Error != nil)
}

func (repo *MovieMySqlRepository) Find(id int64) (models.Movie, error) {
	var entity entities.MovieEntity
	if repo.Db.First(&entity, id).Error != nil {
		return models.Movie{}, fmt.Errorf(fmt.Sprintf("movie with id: %d not found", id))
	}
	movie := mappers.MovieEntityToMovie(entity)
	return movie, nil
}

func (repo *MovieMySqlRepository) FindAll() ([]models.Movie, error) {
	var entities []entities.MovieEntity
	if repo.Db.Find(&entities).Error != nil {
		return nil, fmt.Errorf("no movies found")
	}
	if len(entities) <= 0 {
		return nil, fmt.Errorf("no users found")
	}
	movies := mappers.MoviesEntitiesToMovies(entities)
	return movies, nil
}

func (repo *MovieMySqlRepository) Update(id int64, movie models.Movie) error {
	var current entities.MovieEntity
	if repo.Db.First(&current, id).RecordNotFound() {
		return fmt.Errorf(fmt.Sprintf("error when updated movie to search with id: %v", id))
	}
	if repo.Db.Model(&current).Update(entities.MovieEntity{Name: movie.Name, Director: movie.Director, Writer: movie.Writer, Stars: movie.Stars}).Error != nil {
		return fmt.Errorf(fmt.Sprintf("error when updated movie with id: %v", id))
	}
	return nil
}

func (repo *MovieMySqlRepository) Delete(id int64) error {
	var current entities.MovieEntity
	current.Id = id
	if repo.Db.Delete(current).Error != nil {
		return fmt.Errorf(fmt.Sprintf("cannot delete movie %v", id))
	}
	return nil
}
