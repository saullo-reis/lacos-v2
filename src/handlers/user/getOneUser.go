package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func GetOneUser(c *gin.Context){
	idUser := c.Param("idUser")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error ao conectar com o banco de dados: "+err.Error(),
		})
		return
	}
	defer db.Close()

	var User Users
	row := db.QueryRow("SELECT id_user, username FROM users WHERE id_user = $1", idUser)
	err = row.Scan(&User.Id_user, &User.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"message": "Não encontrou nenhum usuário com esse ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error ao conectar com o banco de dados: "+err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": User,
	})
}