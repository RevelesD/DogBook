package DogControllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type DogsAPI struct {
	Message  []string  `json:"message"`
	Status   string    `json:"status"`
}

func GetRandomDogs() gin.HandlerFunc  {
	fn := func(c *gin.Context) {
		url := "https://dog.ceo/api/breeds/image/random/20"
		response, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusBadGateway, "The HTTP request failed with error " + err.Error() + "\n")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var obj DogsAPI
			if err := json.Unmarshal(data, &obj); err != nil {
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(http.StatusOK, obj)
		}
	}
	return fn
}
