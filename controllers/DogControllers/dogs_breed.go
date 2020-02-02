package DogControllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetDogsByBreed() gin.HandlerFunc  {
	fn := func(c *gin.Context) {
		breed := c.Param("breed")
		sub := c.Param("sub")

		var url string
		if sub == "null" {
			url = "https://dog.ceo/api/breed/" + breed + "/images/random/20"
		} else {
			url = "https://dog.ceo/api/breed/" + breed + "/" + sub + "/images/random/20"
		}
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
