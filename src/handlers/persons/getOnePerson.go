package persons

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)


func GetOnePerson(c * gin.Context){
	idUser := c.Param("idUser")
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	} 

	query := "SELECT person.id_person, person.name, person.birth_date, person.rg, person.cpf, person.cad_unico, person.nis, person.school, person.address, person.address_number, person.blood_type, person.neighborhood, person.city, person.cep, person.home_phone, person.cell_phone, person.contact_phone, person.email, person.current_age, person.active, rperson.id_responsible, rperson.name as rname, rperson.relationship, rperson.rg as rrg, rperson.cpf as rcpf, rperson.cell_phone as rcell_phone FROM persons person LEFT JOIN responsible_person rperson ON person.id_person = rperson.id_person WHERE person.id_person = $1"
	
	var searchPerson SearchFieldsPerson
	err = db.QueryRow(query, idUser).Scan(&searchPerson.IDPerson, &searchPerson.Name, &searchPerson.BirthDate, &searchPerson.RG, &searchPerson.CPF, &searchPerson.CadUnico, &searchPerson.NIS, &searchPerson.School, &searchPerson.Address, &searchPerson.AddressNumber, &searchPerson.BloodType, &searchPerson.Neighborhood, &searchPerson.City, &searchPerson.CEP, &searchPerson.HomePhone , &searchPerson.CellPhone, &searchPerson.ContactPhone, &searchPerson.Email, &searchPerson.CurrentAge, &searchPerson.Active, &searchPerson.ResponsiblePerson.IDResponsible, &searchPerson.ResponsiblePerson.Name, &searchPerson.ResponsiblePerson.Relationship, &searchPerson.ResponsiblePerson.RG, &searchPerson.ResponsiblePerson.CPF, &searchPerson.ResponsiblePerson.CellPhone)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"message": "Nenhuma pessoa encontrada com esse ID",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Error interno no servidor: "+err.Error(),
		})
		return
	}

	responsePerson := creatingResponse(searchPerson)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": responsePerson,
	})
}