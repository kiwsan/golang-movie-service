package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiwsan/golang-movie-service/cmd/application/commands"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/pkg/apierrors"
)

type MovieUpdateController struct {
	IMovieUpdateCommand commands.IMovieUpdateCommand
}

func (movieUpdateController *MovieUpdateController) Put(context *gin.Context) {

	id, movie := movieUpdateController.mapToMovieUpdate(context)
	err := movieUpdateController.IMovieUpdateCommand.Handler(id, movie)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusNoContent, fmt.Sprintf(MovieUpdatedMsg, id))

}

// MovieUpdateController godoc
// @Summary Update a movie
// @Description Update by json movie
// @Tags movies
// @Accept  json
// @Produce  json
// @Param  id path int true "movie id"
// @Param  movie body models.Movie true "Update movie"
// @Success 200 {object} models.Movie
// @Failure 400 {object} apierrors.ApiError
// @Failure 404 {object} apierrors.ApiError
// @Failure 500 {object} apierrors.ApiError
// @Router /movies/{id} [patch]
func (controller *MovieUpdateController) mapToMovieUpdate(c *gin.Context) (id int64, movie models.Movie) {
	id, paramErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if paramErr != nil {
		err := apierrors.NewApiError("Invalid id", paramErr.Error(), 400, nil)
		c.JSON(err.Status(), err)
		return
	}

	var response models.Movie
	if err := c.ShouldBindJSON(&response); err != nil {
		exp := apierrors.NewApiError("Invalid json", err.Error(), 400, nil)
		c.JSON(exp.Status(), exp)
		return
	}
	return id, response
}
