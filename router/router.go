package router

import (
	"github.com/gin-gonic/gin"
	"golang-CICD/handler"
	"golang-CICD/middleware"
	"net/http"
	"os"
)

func SetupServer() *gin.Engine {
	var needMiddleware []gin.HandlerFunc
	runMode := gin.DebugMode
	if os.Getenv("PROGRAM_ENV") == "pro" {
		runMode = gin.ReleaseMode
		needMiddleware = append(needMiddleware, middleware.LogMiddleware())
	}
	gin.SetMode(runMode)
	router := gin.Default()
	for _, handlerFunc := range needMiddleware {
		router.Use(handlerFunc)
	}
	router.HandleMethodNotAllowed = true
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, handler.Response{})
	})
	train := router.Group("/api")
	{
		train.GET("/test", handler.TestHandler)
	}

	return router
}
