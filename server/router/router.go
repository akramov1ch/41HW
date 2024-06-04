package router

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/request", handlers.RequestHandler)
	return r
}
