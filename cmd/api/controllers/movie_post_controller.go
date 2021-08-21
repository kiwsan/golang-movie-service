package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiwsan/golang-movie-service/cmd/application/commands"
	"github.com/kiwsan/golang-movie-service/cmd/domain/models"
	"github.com/kiwsan/golang-movie-service/pkg/apierrors"
)

type MovieCreateController struct {
	IMovieCreateCommand commands.IMovieCreateCommand
}

// MovieCreateController godoc
// @Summary Add a movie
// @Description add by json movie
// @Tags movies
// @Accept  json
// @Produce  json
// @Param movie body models.Movie true "Add movie"
// @Success 200 {object} models.Movie
// @Failure 400 {object} apierrors.ApiError
// @Failure 404 {object} apierrors.ApiError
// @Failure 500 {object} apierrors.ApiError
// @Router /movies [post]
func (controller *MovieCreateController) Post(context *gin.Context) {

	movie := controller.mapToMovie(context)
	err := controller.IMovieCreateCommand.Handler(movie)
	if err != nil {
		abort(context, err)
		return
	}

	context.JSON(http.StatusCreated, fmt.Sprintf(MovieCreatedMsg, movie.Name))

}

func (controller *MovieCreateController) mapToMovie(c *gin.Context) (movie models.Movie) {

	var response models.Movie
	if err := c.ShouldBindJSON(&response); err != nil {
		restErr := apierrors.NewApiError("Invalid json", err.Error(), 400, nil)
		c.JSON(restErr.Status(), restErr)
		return
	}
	return response
}

func abort(ctx *gin.Context, err error) {
	ctx.Error(err)
	ctx.Abort()
}
