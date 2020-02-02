package models

import (
	"errors"
	"fmt"
	"github.com/RevelesD/DogBook/libs/auth"
	"github.com/RevelesD/DogBook/services/mongodb"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"os"
	"time"
)

type AuthDocument struct {
	Email     string			  `json:"email"     bson:"email"`
	Password  string			  `json:"password"  bson:"password"`
	ID        primitive.ObjectID  `json:"_id"       bson:"_id"`
	UserID    primitive.ObjectID  `json:"user_id"   bson:"user_id"`
}

type AuthCreateData struct {
	Email     string			  `json:"email"     bson:"email"`
	Password  string			  `json:"password"  bson:"password"`
	UserID    primitive.ObjectID  `json:"user_id"   bson:"user_id"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type AuthModel struct {
	Client *mongo.Client
}
// Login existing users
func (a AuthModel) SignIn(email string, pass string) (string, error) {
	col, err := mongodb.OpenCollection(a.Client, os.Getenv("DB_NAME"), "Auth")
	if err != nil {
		return "", err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var userFound AuthDocument
	err = col.FindOne(ctx, bson.M{"email": email}).Decode(&userFound)

	if err != nil {
		return "", errors.New("invalid credentials")
	}
	isAuth := auth.ComparePasswords(userFound.Password, pass)
	if !isAuth {
		return "", errors.New("invalid credentials")
	}
	token, err := CreateToken(&userFound.UserID)
	if err != nil {
		return "", err
	}
	return token, nil
}
// Register new user and login the user
func (a AuthModel) SignUp(data *AuthCreateData) (string, error) {
	col, err := mongodb.OpenCollection(a.Client, os.Getenv("DB_NAME"), "Auth")
	if err != nil {
		return "", err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := col.InsertOne(ctx, *data)
	if err != nil {
		fmt.Println(result)
		return "", err
	}
	//var resultOid primitive.ObjectID
	//if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	//	resultOid = oid
	//}

	token, err := CreateToken(&data.UserID)
	if err != nil {
		return "", err
	}
	return token, nil
}
// Returns the count of deleted documents
func (a AuthModel) DeleteEntry(id *primitive.ObjectID) (int64, error) {
	col, err := mongodb.OpenCollection(a.Client, os.Getenv("DB_NAME"), "Auth")
	if err != nil {
		return 0, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := col.DeleteOne(ctx, bson.M{"user_id": *id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func CreateToken(userID *primitive.ObjectID) (string, error)  {
	var claims Claims
	expirationTime := time.Now().Add(2 * time.Hour)
	claims.StandardClaims.ExpiresAt = expirationTime.Unix()
	claims.UserID = userID.Hex()
	fmt.Println(claims.UserID)
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
