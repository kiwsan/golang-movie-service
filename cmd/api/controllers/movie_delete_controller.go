package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiwsan/golang-movie-service/cmd/application/commands"
	"github.com/kiwsan/golang-movie-service/pkg/apierrors"
)

type MovieDeleteController struct {
	IMovieDeleteCommand commands.IMovieDeleteCommand
}

// MovieDeleteController godoc
// @Summary Delete a movie
// @Description Delete by movie id
// @Tags movies
// @Accept  json
// @Produce  json
// @Param  id path int true "movie id" Format(int64)
// @Success 204 {object} models.Movie
// @Failure 400 {object} apierrors.ApiError
// @Failure 404 {object} apierrors.ApiError
// @Failure 500 {object} apierrors.ApiError
// @Router /movies/{id} [delete]
func (controller *MovieDeleteController) Delete(context *gin.Context) {

	id := controller.mapToMovieDelete(context)
	err := controller.IMovieDeleteCommand.Handler(id)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, fmt.Sprintf(MovieDeletedMsg, id))

}

func (controller *MovieDeleteController) mapToMovieDelete(c *gin.Context) (id int64) {
	id, paramErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if paramErr != nil {
		idErr := apierrors.NewApiError("Invalid id", paramErr.Error(), 400, nil)
		c.JSON(idErr.Status(), idErr)
		return
	}
	return id
}
