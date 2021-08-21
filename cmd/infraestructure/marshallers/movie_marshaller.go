package marshallers

import (
	"encoding/json"
	"fmt"

	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
)

type MovieJson struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Stars    string `json:"stars"`
}

func Marshall(movie models.Movie) interface{} {
	movieJson, errUn := json.Marshal(movie)
	fmt.Println(errUn)

	var movieM MovieJson
	_ = json.Unmarshal(movieJson, &movieM)
	return movieM
}

func MarshallArray(movies []models.Movie) []interface{} {
	result := make([]interface{}, len(movies))
	for index, movie := range movies {
		result[index] = Marshall(movie)
	}
	return result
}
