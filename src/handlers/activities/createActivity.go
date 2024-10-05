package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	structs "lacosv2.com/src/handlers/activities/struct"
)

func CreateActivity(c *gin.Context) {
	var body structs.Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "JSON invalido",
		})
		return
	}

	if body.NameActivity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Um nome é obrigatório",
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error interno ao conectar com o banco de dados " +  err.Error(),
		})
		return
	}
	defer db.Close()
	
	_, err = db.Exec("INSERT INTO activity_list(name) VALUES($1) ", body.NameActivity)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro ao inserir a atividade "+ err.Error(),
		})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Atividade inserida com sucesso",
	})
	
}
