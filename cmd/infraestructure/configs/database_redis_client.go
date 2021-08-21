package configs

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
)

const (
	RedisHost   = "REDIS_HOST"
	RedisPort   = "REDIS_PORT"
	RedisSchema = "REDIS_SCHEMA"
	RedisExpire = "REDIS_EXPIRE"
)

func GetClient() *redis.Client {
	address := os.Getenv(RedisHost)
	port := os.Getenv(RedisPort)
	schema, _ := strconv.Atoi(os.Getenv(RedisSchema))

	db := redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: "",
		DB:       schema,
	})

	_, err := db.Ping().Result()
	if err != nil {
		logger.Error(err.Error(), err)
		_ = db.Close()
		panic("Database connection failed")
	}

	return db
}
