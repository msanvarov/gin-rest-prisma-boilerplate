package router

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/controllers"
	_ "github.com/msanvarov/gin-rest-prisma-boilerplate/docs"
	"github.com/msanvarov/gin-rest-prisma-boilerplate/utils"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type RedisConfig struct {
	redisAddress    string
	idleConnections int
	networkType     string
	password        string
	secretKey       []byte
}

// Router method is responsible for binding api routes to methods implemented in the controller.
func Router(config *viper.Viper) *gin.Engine {
	// default gin configuration
	router := gin.Default()
	appEnvironment(config.GetString("server.env"))

	// redis
	store, storeError := redisConnection(&RedisConfig{
		redisAddress:    config.GetString("redis.address"),
		idleConnections: config.GetInt("redis.idle_connections"),
		networkType:     config.GetString("redis.network_type"),
		password:        config.GetString("redis.password"),
		secretKey:       []byte(config.GetString("redis.secret_key")),
	})
	if storeError != nil {
		log.Fatal("Failed to connect to Redis Store.")
	}
	router.Use(sessions.Sessions(config.GetString("session.name"), store))

	casbinConfiguration(router)

	// controllers
	authController := new(controllers.AuthenticationController)
	healthCheckController := new(controllers.HealthCheckController)

	appEndpoint := config.GetString("server.endpoint") + config.GetString("server.port")

	// open api
	url := ginSwagger.URL(fmt.Sprintf("%s/api/v1/spec/doc.json", appEndpoint)) // The url pointing to API definition
	router.GET("/api/v1/spec/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("API specifications are located in %s/api/v1/spec/index.html", appEndpoint))
	})
	router.GET("/ping", healthCheckController.Status)

	api := router.Group("api/v1")
	{
		api.GET("/session", authController.GetSessionData)
		api.POST("/logout", authController.Logout)
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}
	return router
}

func appEnvironment(appEnv string) {
	switch appEnv {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func casbinConfiguration(router *gin.Engine) {
	// casbin
	enforcer := casbin.NewEnforcer(basepath+"/../model.conf", basepath+"/../policy.csv")
	router.Use(func(context *gin.Context) {
		authorizer := &utils.BasicAuthorizer{Enforcer: enforcer}
		if !authorizer.CheckPermission(context) {
			authorizer.RequirePermission(context)
		}
	})
}

func redisConnection(redisConfig *RedisConfig) (redis.Store, error) {
	// session with redis
	return redis.NewStore(redisConfig.idleConnections,
		redisConfig.networkType, redisConfig.redisAddress,
		redisConfig.password, redisConfig.secretKey)
}
