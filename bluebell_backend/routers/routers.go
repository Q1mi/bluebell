package routers

import (
	"bluebell_backend/controller"
	"bluebell_backend/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/auth", controller.JWTHandler)
	r.POST("/login", controller.LoginHandler)
	r.POST("/register", controller.RegisterHandler)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/", controller.IndexHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
