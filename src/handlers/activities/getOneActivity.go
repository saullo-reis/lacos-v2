package activity

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	structs "lacosv2.com/src/handlers/activities/struct"
)

func GetOneActivity(c *gin.Context){
	idActivity := c.Param("idActivity")
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	} 

	query := "SELECT activity.id_activity, activity.name FROM activity_list activity WHERE id_activity = $1"
	
	var searchActivity structs.BodyResponse
	err = db.QueryRow(query, idActivity).Scan(&searchActivity.IdActivity, &searchActivity.NameActivity)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Nenhuma pessoa encontrada com esse ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": searchActivity,
	})
}