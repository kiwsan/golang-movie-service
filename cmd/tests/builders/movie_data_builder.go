package builders

import (
	"fmt"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
)

type MovieDataBuilder struct {
	Id       int64
	Name     string
	Director string
	Writer   string
	Stars    string
}

func NewMovieDataBuilder() *MovieDataBuilder {
	return &MovieDataBuilder{
		Id:       1,
		Name:     "Interstellar",
		Director: "John Doe",
		Writer:   "Jane Doe",
		Stars:    "John Jr Doe, Jane M Doe",
	}
}
func (builder *MovieDataBuilder) Build() models.Movie {
	return models.Movie{
		Id:       builder.Id,
		Name:     builder.Name,
		Director: builder.Director,
		Writer:   builder.Writer,
		Stars:    builder.Stars,
	}
}

func (builder *MovieDataBuilder) BuildString() string {
	return fmt.Sprintf(
		"{\"name\":\"%s\",\"director\":\"%s\",\"writer\":\"%s\",\"stars\":\"%s\"}",
		builder.Name, builder.Director, builder.Writer, builder.Stars,
	)

}
