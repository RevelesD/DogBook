package UserControllers

import (
	"github.com/RevelesD/DogBook/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userModel := models.GetUserModel()
		id := c.Param("id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Error. Provided id is not valid.")
		}
		user, err := userModel.FindOne(&oid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error. Internal Model error.")
		}

		c.JSON(http.StatusOK, user)
	}
	return fn
}