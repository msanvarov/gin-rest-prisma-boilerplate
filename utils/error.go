package utils

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func CreateError(c *gin.Context, status int, err error) {
	payload := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, payload)
}
