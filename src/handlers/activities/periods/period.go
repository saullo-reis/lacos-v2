package period

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

type Period struct {
	IDPeriod int    `json:"id_period"`
	Name     string `json:"name"`
}

func GetPeriods(c *gin.Context){
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error ao conectar com o banco de dados "+ err.Error(),
		})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_period, name FROM period")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao consultar a tabela 'period'",
		})
		return
	}
	defer rows.Close()

	var periods []Period
	for rows.Next() {
		var period Period
		if err := rows.Scan(&period.IDPeriod, &period.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao ler os dados do banco",
			})
			return
		}
		periods = append(periods, period)
	}

	c.JSON(http.StatusOK, periods)
}