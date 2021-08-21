package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiwsan/golang-movie-service/cmd/application/queries"
	"github.com/kiwsan/golang-movie-service/cmd/infraestructure/marshallers"
	"github.com/kiwsan/golang-movie-service/pkg/apierrors"
)

type MovieFindAllController struct {
	IMovieFindAllQuery queries.IMovieFindAllQuery
}

// MovieFindAllController godoc
// @Summary List movies
// @Description get movies
// @Tags movies
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Movie
// @Failure 400 {object} apierrors.ApiError
// @Failure 404 {object} apierrors.ApiError
// @Failure 500 {object} apierrors.ApiError
// @Router /movies [get]
func (controller *MovieFindAllController) Browse(context *gin.Context) {

	movies, err := controller.IMovieFindAllQuery.Handler()

	if err != nil {
		err := apierrors.NewNotFoundApiError(failedFindAllMovies)
		context.JSON(err.Status(), err)
	}

	context.JSON(http.StatusOK, marshallers.MarshallArray(movies))
}
