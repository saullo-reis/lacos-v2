package persons

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func ActivePerson(c *gin.Context){
	idPerson := c.Param("idPerson")

	db, err := database.ConnectDB()
	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro ao conectar com o banco de dados "+ err.Error(),
		})
		return

	}
	var idToVerifyIfExistPerson string
	row := db.QueryRow("SELECT id_person FROM persons WHERE id_person = $1", idPerson)
	err = row.Scan(&idToVerifyIfExistPerson)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Não existe uma pessoa com esse ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro ao verificar pessoa "+ err.Error(),
		})
		return
	}

	_, err = db.Exec("UPDATE persons SET active = 'Y' WHERE id_person = $1", idPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro ao ativar pessoa "+ err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Pessoa ativada com sucesso",
	})

}