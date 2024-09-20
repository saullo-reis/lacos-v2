package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lacosv2.com/src/database/migrations"
	"lacosv2.com/src/handlers/auth"
	"lacosv2.com/src/handlers/persons"
	"lacosv2.com/src/handlers/user"
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

	//AUTH
	r.POST("/login", auth.Login)
	r.POST("/register", auth.AuthMiddlewareAdmin(), auth.Register)

	//USERS
	r.DELETE("/user/delete/:idUser", auth.AuthMiddlewareAdmin(), user.DeleteUser)
	r.GET("/user/get", auth.AuthMiddlewareAdmin(), user.GetAllUsers)
	r.GET("/user/get/:idUser", auth.AuthMiddlewareAdmin(), user.GetOneUser)
	r.PATCH("/user/update/:idUser", auth.AuthMiddlewareAdmin(), user.UpdateUser)

	//PERSONS
	r.POST("/persons/create", auth.AuthMiddleware(), persons.CreatePerson)
	r.Run()
}
