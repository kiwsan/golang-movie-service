package repositories

import "github.com/kiwsan/golang-movie-service/cmd/domain/models"

type MovieRedisRepository interface {
	Set(models.Movie) error
	Get(int64) (models.Movie, error)
}
