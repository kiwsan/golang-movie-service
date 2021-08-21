package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	exp "github.com/kiwsan/golang-movie-service/cmd/domain/exceptions"
	expInternal "github.com/kiwsan/golang-movie-service/cmd/infraestructure/exceptions"
	"github.com/kiwsan/golang-movie-service/pkg/apierrors"
	"github.com/kiwsan/golang-movie-service/pkg/logger"
	"github.com/pkg/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		cause := errors.Cause(err.Err)

		if _, ok := cause.(exp.NotFound); ok {
			throwException(c, http.StatusNotFound, err.Err, cause)
			return
		}

		if _, ok := cause.(exp.Duplicity); ok {
			throwException(c, http.StatusBadRequest, err.Err, cause)
			return
		}

		if _, ok := cause.(expInternal.InternalServerErrorPort); ok {
			throwException(c, http.StatusInternalServerError, err.Err, cause)
			return
		}

		logger.Error("middleware error 500", cause)
		throwException(c, http.StatusInternalServerError, errors.New(internalServerErrorMessage), cause)
	}
}

func throwException(ctx *gin.Context, status int, err error, cause error) {
	restErr := apierrors.NewApiError(err.Error(), http.StatusText(status), status, apierrors.CauseList{cause})
	logger.Error(restErr.Message(), cause)
	ctx.JSON(restErr.Status(), restErr)
}
