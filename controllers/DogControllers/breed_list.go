package DogControllers

import (
	"github.com/RevelesD/DogBook/services/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"time"
)

type BreedList struct {
	List []Breed `json:"list" bson:"list"`
}

type Breed struct {
	Name string `json:"name" bson:"name"`
	Path string `json:"path" bson:"path"`
}

func GetBreedList() gin.HandlerFunc  {
	fn := func(c *gin.Context) {

		con, err := mongodb.GetConnection(os.Getenv("DB_URI"))
		if err != nil {
			log.Fatal("Error. Opening connection with DB", err)
		}
		defer con.Disconnect(context.TODO())

		col, err := mongodb.OpenCollection(con, os.Getenv("DB_NAME"), "Breeds")
		if err != nil {
			log.Fatal("Error. Opening connection with DB", err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		cur, err := col.Find(ctx, bson.D{{}})
		if err != nil {
			log.Fatal("Error. Connection with DB", err)
		}
		defer cur.Close(context.Background())

		var list []Breed

		for cur.Next(context.Background()) {
			var elem Breed
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal("Error. Decoding breed", err)
			}

			list = append(list, elem)
		}

		c.JSON(http.StatusOK, list)
	}
	return fn
}
