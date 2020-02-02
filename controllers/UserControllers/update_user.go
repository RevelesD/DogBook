package UserControllers

import (
	"github.com/RevelesD/DogBook/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func UpdateUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userModel := models.GetUserModel()
		var input struct {
			ID    string                  `json:"id"`
			Data  *models.UserUpdateData  `json:"data"`
		}
		err := c.BindJSON(&input)
		if err != nil {
			log.Fatal("Error. Parameters received doesn't match.", err)
		}
		oid, err := primitive.ObjectIDFromHex(input.ID)
		if err != nil {
			log.Fatal("Error. Provided id is invalid.", err)
		}
		res, err := userModel.Update(&oid, input.Data)
		if err != nil {
			log.Fatal("Error. Internal model error.", err)
		}
		c.JSON(http.StatusOK, res)
	}
	return fn
}