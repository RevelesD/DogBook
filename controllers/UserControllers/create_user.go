package UserControllers

import (
	"github.com/RevelesD/DogBook/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userModel := models.GetUserModel()

		var userData models.UserCreateData
		err := c.BindJSON(&userData)
		if err != nil {
			log.Fatal("Error. Mismatched parameters", err)
		}

		newUser, err := userModel.Create(&userData)
		if err != nil {
			log.Fatal("Error. Internal model error", err)
		}

		c.JSON(http.StatusOK, newUser)
	}
	return fn
}
