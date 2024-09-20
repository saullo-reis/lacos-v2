package persons

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	database "lacosv2.com/src/database/config"
)

type fieldsMandatories struct {
	name   string
	number int
}

func verifyMandatoryField(fields []string) (string, bool) {
	var messageError []string

	fieldsMandatoriesArray := []fieldsMandatories{
		{name: "CPF é obrigatório", number: 1},
		{name: "Nome é obrigatório", number: 2},
		{name: "Data de nascimento é obrigatório", number: 3},
	}

	for i := 0; i < len(fieldsMandatoriesArray); i++ {
		if fields[i] == "" {
			messageError = append(messageError, fieldsMandatoriesArray[i].name)
		}
	}
	returnBoolean := len(messageError) > 0
	return strings.Join(messageError, ", "), returnBoolean
}

func CreatePerson(c *gin.Context) {
	var payloadToCreatePerson PersonJSON
	if err := c.ShouldBindJSON(&payloadToCreatePerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Json inválido",
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Erro interno do servidor: " + err.Error(),
		})
		return
	}
	defer db.Close()

	fieldsMandatories := []string{payloadToCreatePerson.CPF, payloadToCreatePerson.Name, payloadToCreatePerson.BirthDate}
	messageMandatoryVerification, existError := verifyMandatoryField(fieldsMandatories)
	if existError {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": messageMandatoryVerification,
		})
		return
	}

	queryToVerifyDuplicatePerson := "SELECT name FROM persons WHERE CPF = $1"
	var namePersonDuplicated string
	row := db.QueryRow(queryToVerifyDuplicatePerson, payloadToCreatePerson.CPF)
	row.Scan(&namePersonDuplicated)
	
	if namePersonDuplicated != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Já existe uma pessoa com esse CPF",
		})
		return
	}

	var personID int
	err = db.QueryRow(`INSERT INTO persons (
		name, birth_date, rg, cpf, cad_unico, nis, school, address, address_number, blood_type, neighborhood, city, cep, home_phone, cell_phone, contact_phone, email, current_age
	, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18,  'Y') RETURNING id_person`,
		payloadToCreatePerson.Name, payloadToCreatePerson.BirthDate, payloadToCreatePerson.RG, payloadToCreatePerson.CPF, payloadToCreatePerson.CadUnico, payloadToCreatePerson.NIS, payloadToCreatePerson.School, payloadToCreatePerson.Address, payloadToCreatePerson.AddressNumber, payloadToCreatePerson.BloodType, payloadToCreatePerson.Neighborhood, payloadToCreatePerson.City, payloadToCreatePerson.CEP, payloadToCreatePerson.HomePhone, payloadToCreatePerson.CellPhone, payloadToCreatePerson.ContactPhone, payloadToCreatePerson.Email, payloadToCreatePerson.CurrentAge).Scan(&personID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 500,
			"message":     "Erro ao inserir a pessoa no banco de dados",
		})
		return
	}

	if payloadToCreatePerson.ResponsiblePerson == (ResponsiblePersonJSON{}) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"message":     "Pessoa inserida",
		})
		return
	} else {
		_, err = db.Exec(`INSERT INTO responsible_person (
			name, id_person, rg, cpf, relationship, cell_phone) VALUES ($1, $2, $3, $4, $5, $6)`,
			payloadToCreatePerson.ResponsiblePerson.Name, personID, payloadToCreatePerson.ResponsiblePerson.RG, payloadToCreatePerson.ResponsiblePerson.CPF, payloadToCreatePerson.ResponsiblePerson.Relationship, payloadToCreatePerson.ResponsiblePerson.CellPhone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 500,
				"message":     "Erro ao inserir a pessoa responsável no banco de dados",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"message":     "Pessoa registrada com sucesso",
			"data":        payloadToCreatePerson,
		})
	}

}
