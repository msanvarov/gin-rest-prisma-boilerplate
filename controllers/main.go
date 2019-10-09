package auth

import "github.com/gin-gonic/gin"

func GlobalErrorHandler(c *gin.Context, code int, message string) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}