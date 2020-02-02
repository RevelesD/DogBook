package UserControllers

import (
	"github.com/RevelesD/DogBook/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func DeleteAccount() gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var idStruct struct{
			ID string `json:"id"`
		}
		err := c.BindJSON(&idStruct)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		authModel := models.GetAuthModel()
		userModel := models.GetUserModel()
		oid, err := primitive.ObjectIDFromHex(idStruct.ID)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, err)
		}
		authCount, err := authModel.DeleteEntry(&oid)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNoContent, err)
			} else {
				c.JSON(http.StatusInternalServerError, err)
			}
		}
		userCount, err := userModel.Delete(&oid)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNoContent, err)
			} else {
				c.JSON(http.StatusInternalServerError, err)
			}
		}
		var response struct {
			Message string `json:"message"`
		}
		response.Message = "Entry successfully deleted"

		if authCount > 0 && userCount > 0 {
			c.JSON(http.StatusOK, response)
		}
	}
	return fn
}
