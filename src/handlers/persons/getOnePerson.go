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

	var searchPerson SearchFieldsPerson
	err = db.QueryRow("SELECT person.id_person, person.name, person.birth_date, person.rg, person.cpf, person.cad_unico, person.nis, person.school, person.address, person.address_number, person.blood_type, person.neighborhood, person.city, person.cep, person.home_phone, person.cell_phone, person.contact_phone, person.email, person.current_age, person.active, rperson.id_responsible, rperson.name as rname, rperson.relationship, rperson.rg as rrg, rperson.cpf as rcpf, rperson.cell_phone as rcell_phone FROM persons person LEFT JOIN responsible_person rperson ON person.id_person = rperson.id_person WHERE person.id_person = $1", idUser).Scan(&searchPerson.IDPerson, &searchPerson.Name, &searchPerson.BirthDate, &searchPerson.RG, &searchPerson.CPF, &searchPerson.CadUnico, &searchPerson.NIS, &searchPerson.School, &searchPerson.Address, &searchPerson.AddressNumber, &searchPerson.BloodType, &searchPerson.Neighborhood, &searchPerson.City, &searchPerson.CEP, &searchPerson.HomePhone , &searchPerson.CellPhone, &searchPerson.ContactPhone, &searchPerson.Email, &searchPerson.CurrentAge, &searchPerson.Active, &searchPerson.ResponsiblePerson.IDResponsible, &searchPerson.ResponsiblePerson.Name, &searchPerson.ResponsiblePerson.Relationship, &searchPerson.ResponsiblePerson.RG, &searchPerson.ResponsiblePerson.CPF, &searchPerson.ResponsiblePerson.CellPhone)
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

	responseResponsible := ResponseResponsiblePerson{
		IDPerson: searchPerson.ResponsiblePerson.IDPerson.Int64,
		IDResponsible: searchPerson.ResponsiblePerson.IDResponsible.Int64,
		Name: searchPerson.ResponsiblePerson.Name.String,
		Relationship: searchPerson.ResponsiblePerson.Relationship.String,
		RG: searchPerson.ResponsiblePerson.RG.String,
		CPF: searchPerson.ResponsiblePerson.CPF.String,
		CellPhone: searchPerson.ResponsiblePerson.CellPhone.String,
	}

	responsePerson := Response{
		IDPerson: searchPerson.IDPerson.Int64,
		Name: searchPerson.Name.String,
		BirthDate: searchPerson.BirthDate.String,
		RG: searchPerson.RG.String,
		CPF: searchPerson.CPF.String,
		CadUnico: searchPerson.CadUnico.String,
		NIS: searchPerson.NIS.String,
		School: searchPerson.School.String,
		Address: searchPerson.Address.String,
		AddressNumber: searchPerson.AddressNumber.String,
		BloodType: searchPerson.BloodType.String,
		Neighborhood: searchPerson.Neighborhood.String,
		City: searchPerson.City.String,
		CEP: searchPerson.CEP.String,
		HomePhone: searchPerson.HomePhone.String,
		CellPhone: searchPerson.CellPhone.String,
		ContactPhone: searchPerson.ContactPhone.String,
		Email: searchPerson.Email.String,
		CurrentAge: searchPerson.CurrentAge.Int64,
		Active: searchPerson.Active.String,
		ResponsiblePerson: responseResponsible,
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": responsePerson,
	})
}