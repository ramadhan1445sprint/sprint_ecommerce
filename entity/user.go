package entity

type User struct {
	Id       string `db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}