package persons

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
	stct "lacosv2.com/src/handlers/persons/struct"
	funcs "lacosv2.com/src/handlers/persons/funcs"
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

func GetAllPersons(c *gin.Context){
	limit := c.Query("limit")
	offset := c.Query("offset")
	q := c.Query("q")

	if q == "" {
		q = "2 = 2"
	}

	query := "SELECT person.id_person, person.name, person.birth_date, person.rg, person.cpf, person.cad_unico, person.nis, person.school, person.address, person.address_number, person.blood_type, person.neighborhood, person.city, person.cep, person.home_phone, person.cell_phone, person.contact_phone, person.email, person.current_age, person.active, rperson.id_responsible, rperson.name as rname, rperson.relationship, rperson.rg as rrg, rperson.cpf as rcpf, rperson.cell_phone as rcell_phone FROM persons person LEFT JOIN responsible_person rperson ON person.id_person = rperson.id_person WHERE 1 = 1 AND "
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

	var response []stct.Response
	for rows.Next() {
		var searchPerson stct.SearchFieldsPerson
		err = rows.Scan(&searchPerson.IDPerson, &searchPerson.Name, &searchPerson.BirthDate, &searchPerson.RG, &searchPerson.CPF, &searchPerson.CadUnico, &searchPerson.NIS, &searchPerson.School, &searchPerson.Address, &searchPerson.AddressNumber, &searchPerson.BloodType, &searchPerson.Neighborhood, &searchPerson.City, &searchPerson.CEP, &searchPerson.HomePhone , &searchPerson.CellPhone, &searchPerson.ContactPhone, &searchPerson.Email, &searchPerson.CurrentAge, &searchPerson.Active, &searchPerson.ResponsiblePerson.IDResponsible, &searchPerson.ResponsiblePerson.Name, &searchPerson.ResponsiblePerson.Relationship, &searchPerson.ResponsiblePerson.RG, &searchPerson.ResponsiblePerson.CPF, &searchPerson.ResponsiblePerson.CellPhone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"message": "Erro interno no servidor: "+err.Error(),
			})
			return
		}

		response = append(response, funcs.CreatingResponse(searchPerson))
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Sucesso",
		"data": response,
	})
}