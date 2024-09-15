package auth

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func Register(c *gin.Context){
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"message":"Erro ao conectar com o banco de dados",
		})
		return
	}
	defer db.Close()

	var user payload
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"message":     "JSON Invalido",
		})
		return
	}

	if len(user.Password) < 7 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"message": "A senha deve ter no mínimo 8 chars!",
		})
		return
	}

	rows, err := db.Query("SELECT 'Y' FROM users WHERE username = $1", user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"message": "Error de servidor ao buscar pelo usuário",
		})
		return
	}
	defer rows.Close()

	var result sql.NullString
	if rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 500,
				"message": "Error de servidor ao buscar pelo usuário",
			})
			return
		}
	}

	if result.Valid && result.String == "Y" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"message": "Esse usuário já existe!",
		})
		return
	} else {
		hashedPassword := HasherPassword(user.Password)
		db.Exec("INSERT INTO users(username, password) VALUES($1, $2)", user.Username, hashedPassword)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Usuário "+ user.Username + " criado",
			"status_code": 201,
		})
		return
	}
}