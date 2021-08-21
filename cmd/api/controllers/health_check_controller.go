package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct {
}

// HealthCheckController godoc
// @summary Health Check
// @description Health checking for the service
// @Tags monitoring
// @id HealthCheckController
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func (controller *HealthCheckController) HealthCheck(context *gin.Context) {
	context.String(http.StatusOK, "OK")
}
