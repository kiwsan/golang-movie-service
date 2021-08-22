package repositories_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/configs"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/repositories"
	"github.com/kiwsan/golang-movie-service/cmd/tests/builders"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	movieMysqlRepository repositories.MovieMySqlRepository
)

func TestMain(m *testing.M) {
	containerMockServer, ctx := load()
	containerMockServerR, ctx := loadRedis()
	code := m.Run()
	beforeAll(containerMockServer, ctx)
	beforeAll(containerMockServerR, ctx)
	os.Exit(code)
}

func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "ceiba.2020",
			"MYSQL_DATABASE":      "movies_db",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}
	mysqlC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Port()
	_ = os.Setenv("MYSQL_HOST", host)
	_ = os.Setenv("MYSQL_PORT", port)
	_ = os.Setenv("MYSQL_SCHEMA", "movies_db")
	_ = os.Setenv("MYSQL_USERNAME", "root")
	_ = os.Setenv("MYSQL_PASSWORD", "ceiba.2020")

	movieMysqlRepository = repositories.MovieMySqlRepository{
		Db: configs.GetDatabaseInstance(),
	}
	return mysqlC, ctx
}

func beforeAll(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
func TestMovieMysqlRepository_Create(t *testing.T) {
	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie models.Movie
	movie = builders.NewMovieDataBuilder().Build()
	err := movieMysqlRepository.Create(movie)

	assert.Nil(t, err)
}

func TestMovieMysqlRepository_Find(t *testing.T) {

	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie models.Movie
	movie = builders.NewMovieDataBuilder().Build()
	err := movieMysqlRepository.Create(movie)

	movieFind, err := movieMysqlRepository.Find(movie.Id)

	assert.Nil(t, err)
	assert.NotNil(t, movieFind)
	assert.EqualValues(t, movie.Id, movieFind.Id)
	assert.EqualValues(t, movie.Name, movieFind.Name)
	assert.EqualValues(t, movie.Director, movieFind.Director)
	assert.EqualValues(t, movie.Writer, movieFind.Writer)
	assert.EqualValues(t, movie.Stars, movieFind.Stars)
}

func TestMovieMysqlRepository_FindAll(t *testing.T) {
	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()

	var movie models.Movie
	movie = builders.NewMovieDataBuilder().Build()
	_ = movieMysqlRepository.Create(movie)

	movieFind, err := movieMysqlRepository.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, movieFind)
}

func TestMovieMysqlRepository_Update(t *testing.T) {
	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie models.Movie
	movie = builders.NewMovieDataBuilder().Build()
	_ = movieMysqlRepository.Create(movie)

	movie.Director = "Jane Doe"
	errUpdate := movieMysqlRepository.Update(1, movie)

	assert.Nil(t, errUpdate)
	assert.EqualValues(t, movie.Director, "Jane Doe", "Director names are differences")
	assert.EqualValues(t, movie.Stars, "John Jr Doe, Jane M Doe")
	assert.NotEqual(t, movie.Director, "John Doe")
}

func TestMovieMysqlRepository_Delete(t *testing.T) {
	tx := movieMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var movie models.Movie
	movie = builders.NewMovieDataBuilder().Build()
	_ = movieMysqlRepository.Create(movie)

	errDelete := movieMysqlRepository.Delete(1)

	assert.Nil(t, errDelete)
	_, err := movieMysqlRepository.Find(movie.Id)
	assert.NotNil(t, err)
	assert.EqualError(t, errors.New("movie with id: 1 not found"), err.Error())
}
