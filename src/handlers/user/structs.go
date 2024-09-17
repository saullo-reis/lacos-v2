package user

type Users struct {
	Id_user string `json:"id_user"`
	Username string `json:"username"`
}

type UpdateUsers struct {
	Username string `json:"username"`
	Password string `json:"password"`
}