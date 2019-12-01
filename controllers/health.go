package controllers

import "github.com/gin-gonic/gin"

import "net/http"

// IHealthCheckController interface
type IHealthCheckController interface {
	Status(c *gin.Context)
}

// HealthCheckController for health checking.
type HealthCheckController struct{}

// Status method is responsible for indicating whether or not the server is working.
// @summary Checks status of the server
// @tags health
// @produce text/plain
// @success 200 {string} string	"Status Request Completed"
// @failure 400 {object} utils.HTTPError
// @router /ping [get]
func (HealthCheckController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
