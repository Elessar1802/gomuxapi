package repository

type User struct {
	Id    string `json:"id,notnull"`
	Name  string `json:"name,notnull"`
	Phone string `json:"phone,notnull"`
	Role  string   `json:"role,notnull"`
}
