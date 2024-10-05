package activity

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func DeleteActivity(c *gin.Context){
	idActivity := c.Param("idActivity")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error ao conectar com o banco de dados "+ err.Error(),
		})
		return
	}
	defer db.Close()

	var idToVerifyIfExistActivity string
	row := db.QueryRow("SELECT id_activity FROM activity_list WHERE id_activity = $1", idActivity)
	err = row.Scan(&idToVerifyIfExistActivity)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Não existe uma atividade com esse ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro ao verificar atividade "+ err.Error(),
		})
		return
	}

	_, err = db.Exec("DELETE FROM activity_list WHERE id_activity = $1", idActivity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro ao deletar atividade "+ err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Atividade deletada com sucesso",
	})
}