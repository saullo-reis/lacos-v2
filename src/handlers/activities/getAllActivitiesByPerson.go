package activity

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	structs "lacosv2.com/src/handlers/activities/struct"
)

func GetAllActivitiesByPerson(c *gin.Context){
	idPerson := c.Param("idPerson")

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro interno no servidor err: "+ err.Error(),
		})
		return
	}
	defer db.Close()

	var existPerson string 
	row := db.QueryRow("SELECT 'Y' from persons WHERE id_person = $1", idPerson)
	err = row.Scan(&existPerson)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Esse id de pessoa não existe",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro interno no servidor err: "+ err.Error(),
		})
		return
	}

	query := `SELECT A.id_activity, A.id_period,A.id_person,A.id_activities,B.name,
							C.name,
							A.hour_start,
							A.hour_end
						FROM activities A
						JOIN activity_list B ON A.id_activity = B.id_activity 
						JOIN period C ON A.id_period = C.id_period
						WHERE A.id_person = $1`
	rows, err := db.Query(query, idPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Erro interno no servidor err: "+ err.Error(),
		})
		return
	}
	defer rows.Close()

	var response []structs.ResponseActivitiesByPerson
	for rows.Next() {
		var responseUnique structs.ResponseActivitiesByPerson

		err = rows.Scan(&responseUnique.IdActivity, &responseUnique.IdPeriod, &responseUnique.IdPerson, &responseUnique.IdActivities, &responseUnique.NameActivity, &responseUnique.NamePeriod, &responseUnique.HourStart, &responseUnique.HourEnd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"message": "Erro interno no servidor err: "+ err.Error(),
			})
			return
		}

		response = append(response, responseUnique)
	}

	if len(response) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"message": "Essa pessoa não tem nenhuma atividade",
		})	
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": response,
	})

}