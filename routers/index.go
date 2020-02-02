package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// load groups

	v1 := router.Group("/api/v1")
	LoadUserRoutes(v1, "/user")
	LoadDogsRoutes(v1, "/dog")
	LoadAuthRoutes(v1, "/auth")

	return router
}
