package repositories

import "github.com/kiwsan/golang-movie-service/cmd/domain/models"

type MovieRepository interface {
	Create(models.Movie) error
	Exist(string) bool
	Find(int64) (models.Movie, error)
	FindAll() ([]models.Movie, error)
	Update(int64, models.Movie) error
	Delete(int64) error
}
