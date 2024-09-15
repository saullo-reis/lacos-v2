package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lacosv2.com/src/database/migrations"
	"lacosv2.com/src/handlers/auth"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error ao carregar variáveis de ambiente " + err.Error())
	}
}

func main() {
	r := gin.Default()
	migrations.CreateTables()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "pong",
		})
	})
	r.POST("/login", auth.Login)
	r.POST("/register", auth.AuthMiddlewareAdmin(), auth.Register)

	r.Run()
}
