package persons

import stct "lacosv2.com/src/handlers/persons/struct"

func CreatingResponse(person stct.SearchFieldsPerson) stct.Response{
	responseResponsible := stct.ResponseResponsiblePerson{
		IDPerson: person.ResponsiblePerson.IDPerson.Int64,
		IDResponsible: person.ResponsiblePerson.IDResponsible.Int64,
		Name: person.ResponsiblePerson.Name.String,
		Relationship: person.ResponsiblePerson.Relationship.String,
		RG: person.ResponsiblePerson.RG.String,
		CPF: person.ResponsiblePerson.CPF.String,
		CellPhone: person.ResponsiblePerson.CellPhone.String,
	}

	responsePerson := stct.Response{
		IDPerson: person.IDPerson.Int64,
		Name: person.Name.String,
		BirthDate: person.BirthDate.String,
		RG: person.RG.String,
		CPF: person.CPF.String,
		CadUnico: person.CadUnico.String,
		NIS: person.NIS.String,
		School: person.School.String,
		Address: person.Address.String,
		AddressNumber: person.AddressNumber.String,
		BloodType: person.BloodType.String,
		Neighborhood: person.Neighborhood.String,
		City: person.City.String,
		CEP: person.CEP.String,
		HomePhone: person.HomePhone.String,
		CellPhone: person.CellPhone.String,
		ContactPhone: person.ContactPhone.String,
		Email: person.Email.String,
		CurrentAge: person.CurrentAge.Int64,
		Active: person.Active.String,
		ResponsiblePerson: responseResponsible,
	}

	return responsePerson
}
