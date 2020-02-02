package main

import (
	"github.com/RevelesD/DogBook/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.DebugMode)
	r := routers.SetupRouter()
	r.Run(":8080")
}
