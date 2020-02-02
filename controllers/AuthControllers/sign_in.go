package AuthControllers

import (
	"github.com/RevelesD/DogBook/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var credentials struct{
			Email     string  `json:"email"`
			Password  string  `json:"password"`
		}
		err := c.BindJSON(&credentials)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		authModel := models.GetAuthModel()
		token, err := authModel.SignIn(credentials.Email, credentials.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
		}
		c.JSON(http.StatusOK, token)
	}
	return fn
}
