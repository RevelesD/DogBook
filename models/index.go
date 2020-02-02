package models

import (
	"github.com/RevelesD/DogBook/services/mongodb"
	"log"
	"os"
)

func GetUserModel() UserModel {
	con, err := mongodb.GetConnection(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Error. Opening connection with DB", err)
	}
	userModel := UserModel{Client: con}
	return userModel
}

func GetAuthModel() AuthModel {
	con, err := mongodb.GetConnection(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Error. Opening connection with DB", err)
	}
	authModel := AuthModel{Client: con}
	return authModel
}