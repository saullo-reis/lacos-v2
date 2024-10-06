package activity

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	structs "lacosv2.com/src/handlers/activities/struct"
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

func GetAllActivities(c *gin.Context){
	limit := c.Query("limit")
	offset := c.Query("offset")
	q := c.Query("q")

	if q == "" {
		q = "2 = 2"
	}

	query := "SELECT id_activity, name FROM activity_list WHERE 1 = 1 AND "
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

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro interno no servidor: "+err.Error(),
		})
		return
	}
	defer rows.Close()

	var response []structs.BodyResponse

	for rows.Next() {
		var searchActivity structs.BodyResponse
		err = rows.Scan(&searchActivity.IdActivity, &searchActivity.NameActivity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"message": "Erro interno no servidor: "+err.Error(),
			})
			return
		}

		response = append(response, searchActivity)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": response,
	})
}