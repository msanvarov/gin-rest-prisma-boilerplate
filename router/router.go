package router

import (
	"github.com/gin-gonic/gin"
	auth "github.com/gin-rest-prisma-boilerplate/controllers"
	"net/http"
)

var authController = new(auth.AuthenticationController)

func Router() *gin.Engine {
	// default gin configuration
	router := gin.Default()

	// custom middleware
	router.Use(func(c *gin.Context) {

		// setting cors headers for integration with front end frameworks like angular/react/vue
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Api-Key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	api := router.Group("api")
	{
		api.POST("/login")
		api.POST("/register", authController.Register)
	}
	return router
}