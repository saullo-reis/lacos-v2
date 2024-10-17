package activity

import (
	"net/http"
	"strings"

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

	var motiveMandatory []string
	if body.NameActivity == "" {
		motiveMandatory = append(motiveMandatory, "Nome obrigatório")
	}
	if body.HourStart == "" {
		motiveMandatory = append(motiveMandatory, "Horário de início obrigatório")
	}
	if body.HourEnd == "" {
		motiveMandatory = append(motiveMandatory, "Horário de fim obrigatório")
	}

	if len(motiveMandatory) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": strings.Join(motiveMandatory, ", "),
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
	
	_, err = db.Exec("INSERT INTO activity_list(name, hour_start, hour_end, id_period) VALUES($1, $2, $3, $4) ", body.NameActivity, body.HourStart, body.HourEnd, body.IdPeriod)
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
