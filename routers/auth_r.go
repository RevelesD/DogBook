package routers

import (
	ac "github.com/RevelesD/DogBook/controllers/AuthControllers"
	"github.com/gin-gonic/gin"
)

func LoadAuthRoutes(group *gin.RouterGroup, route string)  {
	auth := group.Group(route)
	auth.POST("/signin", ac.SignIn())
	auth.POST("/signup", ac.SignUp())
}
