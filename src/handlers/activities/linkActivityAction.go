package activity

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	structs "lacosv2.com/src/handlers/activities/struct"
)

func LinkActivity(c *gin.Context){
	var request structs.LinkActivityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	var existPerson string
	row := db.QueryRow("SELECT 'Y' FROM persons WHERE id_person = $1", request.IdPerson)
	err = row.Scan(&existPerson)
	if err != nil{
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Pessoa não existe",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	var existPeriod string
	row = db.QueryRow("SELECT 'Y' FROM period WHERE id_period = $1", request.IdPeriod)
	err = row.Scan(&existPeriod)
	if err != nil{
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Período não existe",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	var existActivity string
	row = db.QueryRow("SELECT 'Y' FROM activity_list WHERE id_activity = $1", request.IdActivity)
	err = row.Scan(&existActivity)
	if err != nil{
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Atividade não existe",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	_, err = db.Exec("INSERT INTO activities(id_person, id_activity, id_period, hour_start, hour_end) VALUES($1, $2, $3, $4, $5)", request.IdPerson, request.IdActivity, request.IdPeriod, request.HourStart, request.HourEnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Atividade linkada com sucesso",
	})
	
}