package routers

import (
	uc "github.com/RevelesD/DogBook/controllers/UserControllers"
	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(group *gin.RouterGroup, route string)  {
	auth := group.Group(route)
	auth.GET("/id/:id", uc.GetUser())
	auth.POST("/createUser", uc.CreateUser())
	auth.POST("/updateUser", uc.UpdateUser())
	auth.POST("/deleteAccount", uc.DeleteAccount())
}
