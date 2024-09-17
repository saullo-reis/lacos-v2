package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func DeleteUser(c *gin.Context){
	idUser := c.Param("idUser")
	db, err := database.ConnectDB();
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}
	defer db.Close()
	
	var username string
	row, err := db.Query("SELECT username FROM users WHERE id_user = $1", idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error ao verificar se usuário existe no banco de dados: "+err.Error(),
		})
		return
	}
	for row.Next() {
		row.Scan(&username)
		db.Exec("DELETE FROM users WHERE id_user = $1", idUser)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"message": "Usuário deletado com sucesso",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": 400,
		"message": "Esse usuário não existe",
	})
}