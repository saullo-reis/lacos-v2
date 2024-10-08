package activity

type Body struct {
	NameActivity string `json:"name"`
}

type BodyResponse struct {
	NameActivity string `json:"name"`
	IdActivity int64 `json:"id_activity"`
}

type LinkActivityRequest struct {
	IdActivity int64 `json:"id_activity"`
	IdPeriod int64 `json:"id_period"`
	IdPerson int64 `json:"id_person"`
	HourStart string `json:"hour_start"`
	HourEnd string `json:"hour_end"`
}