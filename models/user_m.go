package models

import (
	dc "github.com/RevelesD/DogBook/controllers/DogControllers"
	"github.com/RevelesD/DogBook/libs"
	"github.com/RevelesD/DogBook/services/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"os"
	"time"
)

type UserUpdateData struct {
	Name      string   `json:"name,omitempty"      bson:"name,omitempty"`
	Birthday  string   `json:"birthday,omitempty"  bson:"birthday,omitempty"`
	Email     string   `json:"email,omitempty"     bson:"email,omitempty"`
	Breed     dc.Breed `json:"breed,omitempty"     bson:"breed,omitempty"`
}

type UserCreateData struct {
	Name      string   `json:"name"       bson:"name"`
	Birthday  string   `json:"birthday"   bson:"birthday"`
	Email     string   `json:"email"      bson:"email"`
	Breed     dc.Breed `json:"breed"      bson:"breed"`
}

type UserDocument struct {
	ID        *primitive.ObjectID  `json:"_id,omitempty"  bson:"_id,omitempty"`
	Name      string 			   `json:"name"      	  bson:"name"`
	Birthday  string			   `json:"birthday"  	  bson:"birthday"`
	Email     string               `json:"email"     	  bson:"email"`
	Breed     dc.Breed             `json:"breed"     	  bson:"breed"`
}

type UserModel struct {
	Client *mongo.Client
}
// returns a new user with the id just created by the database
func (u UserModel) Create(data *UserCreateData) (*UserDocument, error)  {
	col, err := mongodb.OpenCollection(u.Client, os.Getenv("DB_NAME"), "Users")
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := col.InsertOne(ctx, *data)
	if err != nil {
		return nil, err
	}

	var newDocument UserDocument

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		newDocument.ID = &oid
	}

	newDocument.Breed = data.Breed
	newDocument.Email = data.Email
	newDocument.Birthday = data.Birthday
	newDocument.Name = data.Name

	return &newDocument, nil
}
// Returns the count of deleted documents
func (u UserModel) Delete(id *primitive.ObjectID) (int64, error)  {
	col, err := mongodb.OpenCollection(u.Client, os.Getenv("DB_NAME"), "Users")
	if err != nil {
		return 0, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := col.DeleteOne(ctx, bson.M{"_id": *id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
// Returns updated document
func (u UserModel) Update(id *primitive.ObjectID, data *UserUpdateData) (*UserDocument, error)  {
	col, err := mongodb.OpenCollection(u.Client, os.Getenv("DB_NAME"), "Users")
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var updatedDocument UserDocument
	document, err := libs.StructToBson(*data)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *id}
	update := bson.D{{"$set", document}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	opts.SetUpsert(false)

	err = col.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		return nil, err
	}

	return &updatedDocument, nil
}
// Find by id...
func (u UserModel) FindOne(id *primitive.ObjectID) (*UserDocument, error)  {
	col, err := mongodb.OpenCollection(u.Client, os.Getenv("DB_NAME"), "Users")
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var userFound UserDocument

	err = col.FindOne(ctx, bson.M{"_id": *id}).Decode(&userFound)
	if err != nil {
		return nil, err
	}
	return &userFound, nil
}
