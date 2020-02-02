package routers

import (
	dc "github.com/RevelesD/DogBook/controllers/DogControllers"
	"github.com/gin-gonic/gin"
)

func LoadDogsRoutes(group *gin.RouterGroup, route string)  {
	dog := group.Group(route)
	dog.GET("/getBreedList", dc.GetBreedList())
	dog.GET("/getProfilePic/:breed/:sub", dc.GetProfilePic())
	dog.GET("/getDogsRandom", dc.GetRandomDogs())
	dog.GET("/getDogsBreed/:breed/:sub", dc.GetDogsByBreed())
}
