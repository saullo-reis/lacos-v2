package auth

type Payload struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string
	Password string
}