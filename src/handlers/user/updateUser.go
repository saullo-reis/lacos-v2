package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	"lacosv2.com/src/handlers/auth"
	stct "lacosv2.com/src/handlers/user/struct"
)

func UpdateUser(c *gin.Context) {
	id_user := c.Param("idUser")
	var body stct.UpdateUsers
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "JSON inválido: " + err.Error(),
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error ao conectar com o banco de dados: " + err.Error(),
		})
		return
	}
	defer db.Close()

	var username string
	row := db.QueryRow("SELECT username FROM users WHERE id_user = $1", id_user)
	err = row.Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Usuário não existe",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error ao verificar usuário no banco de dados: " + err.Error(),
		})
		return
	}

	var updateDatas []string
	query := "UPDATE users SET "
	if body.Username != "" {
		updateDatas = append(updateDatas, "username = '" + body.Username + "'")
	}
	if body.Password != "" {
		updateDatas = append(updateDatas, "password = '" + auth.HasherPassword(body.Password) + "'")
	}
	query = query + strings.Join(updateDatas, ", ") + " WHERE id_user = $1"

	fmt.Println(query)
	up, err := db.Exec(query, id_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  500,
			"message": "Erro ao atualizar usuário: " + err.Error(),
		})
		return
	}
	fmt.Println(up)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Usuário atualizado",
		"data":    body,
	})

}
