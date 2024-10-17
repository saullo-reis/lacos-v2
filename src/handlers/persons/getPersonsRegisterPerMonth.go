package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config" 
)

type bodyResponseStruct struct {
    Month string `json:"month"`
    Count int `json:"count"`
}

func GetPersonsRegisteredPerMonth(c *gin.Context){
    query := `SELECT 
                EXTRACT(MONTH FROM created_date) AS mes,
                COUNT(*) AS total_pessoas
                FROM persons
                WHERE EXTRACT(YEAR FROM created_date) = EXTRACT(YEAR FROM CURRENT_DATE)
                GROUP BY mes
                ORDER BY mes;` 
    db, err := database.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
    }

    rows, err := db.Query(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
    }

    var bodyResponse []bodyResponseStruct
    for rows.Next() {
        var row bodyResponseStruct
        err = rows.Scan(&row.Month, &row.Count)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "status": 500,
                "message": "Error interno no servidor: "+err.Error(),
            })
            return
        }

        bodyResponse = append(bodyResponse, row)
    }
    c.JSON(http.StatusOK, gin.H{
        "status": 200,
        "message": "Sucesso",
        "data": bodyResponse,
    })
}