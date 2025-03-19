package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/back-go-barbershop-core/api/handlers"
	"github.com/samuelmaxi/back-go-barbershop-core/backend/config"
)

func main() {
	router := gin.Default()
	var database config.DataBase
	db := database.Connect()

	router.POST("/", handlers.RegisterUser(db))

	router.Run(":8081")
}
