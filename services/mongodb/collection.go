package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func OpenCollection(client *mongo.Client, db string, col string) (*mongo.Collection, error) {
	//client, err := GetConnection(os.Getenv("DB_URI"))
	//if err != nil {
	//	return nil, err
	//}
	collection := client.Database(db).Collection(col)
	return collection, nil
}
