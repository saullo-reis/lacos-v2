package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

func addingLimit(limit string) string {
	if limit == "" {
		return ""
	}
	return "LIMIT " + limit
}

func addingOffset(offset string) string {
	if offset == "" {
		return ""
	}
	return "OFFSET " + offset
}

func GetAllUsers(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	q := c.Query("q")

	if q == "" {
		q = "2 = 2"
	}

	query := "SELECT id_user, username FROM users WHERE 1 = 1 AND "
	query = query + " " + q + " "
	query = query + addingLimit(limit) + " " + addingOffset(offset)
	fmt.Println(query)
	db, err := database.ConnectDB()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error ao conectar com o banco de dados: " + err.Error(),
		})
		return
	}
	defer db.Close()

	var response []Users
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error ao conectar com o banco de dados: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var oneRowScan Users
		rows.Scan(&oneRowScan.Id_user, &oneRowScan.Username)
		response = append(response, oneRowScan)
	}

	if len(response) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Nenhum usuário encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Sucesso",
		"data":    response,
	})

}
