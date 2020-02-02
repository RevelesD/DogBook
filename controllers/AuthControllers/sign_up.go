package AuthControllers

import (
	dc "github.com/RevelesD/DogBook/controllers/DogControllers"
	"github.com/RevelesD/DogBook/libs/auth"
	"github.com/RevelesD/DogBook/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input struct{
			Email     string   `json:"email"      bson:"email"`
			Password  string   `json:"password"   bson:"password"`
			Name      string   `json:"name"       bson:"name"`
			Breed     dc.Breed `json:"breed"      bson:"breed"`
			Birthday  string   `json:"birthday"   bson:"birthday"`
		}
		err := c.BindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		createUser := models.UserCreateData{
			Name:     input.Name,
			Birthday: input.Birthday,
			Email:    input.Email,
			Breed:    input.Breed,
		}
		userModel := models.GetUserModel()
		newUser, err := userModel.Create(&createUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		createAuth := models.AuthCreateData{
			Email:    newUser.Email,
			Password: auth.HashAndSalt(input.Password),
			UserID:   *newUser.ID,
		}
		authModel := models.GetAuthModel()
		token, err := authModel.SignUp(&createAuth)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, token)
	}
	return fn
}