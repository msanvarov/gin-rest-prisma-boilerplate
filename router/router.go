package router

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/controllers"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/utils"
	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// Router method is responsible for binding api routes to methods implemented in the controller.
func Router(config *viper.Viper) *gin.Engine {
	// default gin configuration
	router := gin.Default()

	switch config.GetString("server.env") {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	var redisAddress string

	if redisEnvAddress := os.Getenv("REDIS_CONNECTION"); redisEnvAddress != "" {
		redisAddress = redisEnvAddress
	} else {
		redisAddress = config.GetString("redis.address")
	}

	// session with redis
	for {
		if store, redisErr := redis.NewStore(config.GetInt("redis.idle_connections"),
			config.GetString("redis.network_type"), redisAddress,
			config.GetString("redis.password"), []byte(config.GetString("redis.secret_key"))); redisErr != nil {
			log.Print(redisErr)
			time.Sleep(20 * time.Second)
			continue
		} else {
			router.Use(sessions.Sessions(config.GetString("session.name"), store))
			break
		}
	}

	// casbin
	enforcer := casbin.NewEnforcer(basepath+"/../model.conf", basepath+"/../policy.csv")
	router.Use(func(c *gin.Context) {
		authorizer := &utils.BasicAuthorizer{Enforcer: enforcer}
		if !authorizer.CheckPermission(c) {
			authorizer.RequirePermission(c)
		}
	})

	// controllers
	authController := new(controllers.AuthenticationController)

	api := router.Group("api/v1")
	{
		api.GET("/session", authController.GetSessionData)
		api.GET("/logout", authController.Logout)
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}
	return router
}
