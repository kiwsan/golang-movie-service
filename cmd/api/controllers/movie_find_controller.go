package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiwsan/golang-movie-service/cmd/application/queries"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/marshallers"
	"github.com/kiwsan/golang-movie-service/pkg/apierrors"
)

type MovieFindController struct {
	IMovieFindQuery queries.IMovieFindQuery
}

// MovieFindController godoc
// @Summary Show a movie
// @Description get string by id
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "movie id"
// @Success 200 {object} models.Movie
// @Failure 400 {object} apierrors.ApiError
// @Failure 404 {object} apierrors.ApiError
// @Failure 500 {object} apierrors.ApiError
// @Router /movies/{id} [get]
func (controller *MovieFindController) Find(context *gin.Context) {
	id := controller.mapToMovieFind(context)
	movie, err := controller.IMovieFindQuery.Handler(id)

	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusOK, marshallers.Marshall(movie))

}

func (controller *MovieFindController) mapToMovieFind(c *gin.Context) (id int64) {
	id, exp := strconv.ParseInt(c.Param("id"), 10, 64)
	if exp != nil {
		err := apierrors.NewApiError("Invalid id", exp.Error(), 400, nil)
		c.JSON(err.Status(), err)
		return
	}
	return id
}
