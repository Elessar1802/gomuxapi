package repository

type User struct {
	Id    int    `json:"id,notnull" pg:"type:serial,pk"`
	Name  string `json:"name,notnull" pg:",notnull"`
	Phone string `json:"phone,notnull" pg:",notnull"`
	Role  string `json:"role,notnull,omitempty" pg:",notnull"`
	Class string `json:"class,omitempty" pg:"-"`
}
