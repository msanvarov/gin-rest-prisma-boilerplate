package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-prisma-boilerplate/controllers"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func Router(config *viper.Viper) *gin.Engine {
	// default gin configuration
	router := gin.Default()

	env := config.GetString("server.env")
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// session with redis
	store, storeError := sessions.NewRedisStore(config.GetInt("redis.idle_connections"),
		config.GetString("redis.network_type"), config.GetString("redis.address"),
		config.GetString("redis.password"), []byte(config.GetString("redis.secret_key")))

	if storeError != nil {
		log.Fatal(storeError)
	}

	router.Use(sessions.Sessions(config.GetString("session.name"), store))

	// controllers
	authController := new(controllers.AuthenticationController)

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

	api := router.Group("api/v1")
	{
		api.GET("/session", authController.GetSessionData)
		api.GET("/logout", authController.Logout)
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}
	return router
}
