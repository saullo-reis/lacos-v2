package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lacosv2.com/src/database/migrations"
	activity "lacosv2.com/src/handlers/activities"
	"lacosv2.com/src/handlers/auth"
	"lacosv2.com/src/handlers/persons"
	"lacosv2.com/src/handlers/user"
	"github.com/gin-contrib/cors"
)

// func init() {
// 	err := godotenv.Load("./.env")
// 	if err != nil {
// 		log.Fatalf("Error ao carregar variáveis de ambiente " + err.Error())
// 	}
// }

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, 
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
	r.GET("/persons/get/:idUser", auth.AuthMiddleware(), persons.GetOnePerson)
	r.GET("/persons/get", auth.AuthMiddleware(), persons.GetAllPersons)
	r.PATCH("/persons/update", auth.AuthMiddleware(), persons.UpdatePersons)
	r.DELETE("/persons/delete/:idPerson", auth.AuthMiddleware(), persons.DeletePerson)
	r.POST("/persons/active/:idPerson", auth.AuthMiddleware(), persons.ActivePerson)
	r.GET("/persons/get/monthRegistered", auth.AuthMiddleware(), persons.GetPersonsRegisteredPerMonth)

	//ACTIVITIES
	r.POST("/activityList/create", auth.AuthMiddleware(), activity.CreateActivity)
	r.DELETE("/activityList/delete/:idActivity", auth.AuthMiddleware(), activity.DeleteActivity)
	r.GET("/activityList/get/:idActivity", auth.AuthMiddleware(), activity.GetOneActivity)
	r.GET("/activityList/get", auth.AuthMiddleware(), activity.GetAllActivities)
	r.POST("/activities/action/link", auth.AuthMiddleware(), activity.LinkActivity)
	r.DELETE("/activities/action/link/delete/:idActivities", auth.AuthMiddleware(), activity.DeleteLink)
	r.GET("/activities/getAll/:idPerson", auth.AuthMiddleware(), activity.GetAllActivitiesByPerson)
	r.Run()
}
