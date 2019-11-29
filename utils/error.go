package utils

import "github.com/gin-gonic/gin"

// HTTPError structure to define response payload on error.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// CreateError responds to client requests with error payloads.
func CreateError(c *gin.Context, status int, errMessage string) {
	payload := HTTPError{
		Code:    status,
		Message: errMessage,
	}
	c.JSON(status, payload)
}
