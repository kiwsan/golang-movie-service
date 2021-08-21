package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
)

type MovieRedisRepository struct {
	Db *redis.Client
}

func (repo *MovieRedisRepository) Set(movie models.Movie) error {
	exp, _ := strconv.Atoi(os.Getenv(RedisExpire))
	json, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}

	if err := repo.Db.Set(strconv.FormatInt(movie.Id, 10), json, time.Duration(exp)*time.Second).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *MovieRedisRepository) Get(id int64) (models.Movie, error) {

	val, err := repo.Db.Get(strconv.FormatInt(id, 10)).Result()
	if err != nil {
		return models.Movie{}, fmt.Errorf(fmt.Sprintf("movie with id: %d not found", id))
	}

	movie := models.Movie{}
	err = json.Unmarshal([]byte(val), &movie)
	if err != nil {
		panic(err)
	}
	return movie, nil
}
