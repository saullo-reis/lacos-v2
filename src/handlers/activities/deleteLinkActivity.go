package activity

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func DeleteLink(c *gin.Context){
	idActivities := c.Param("idActivities")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	var existActivities string
	row := db.QueryRow("SELECT 'Y' FROM activities WHERE id_activities = $1", idActivities)
	err = row.Scan(&existActivities)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 400,
				"message": "Esse id de atividade não existe para essa pessoa",
			})
			return

		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	_, err = db.Exec("DELETE FROM activities WHERE id_activities = $1", idActivities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Atividade retirada dessa pessoa com sucesso",
	})
	
}