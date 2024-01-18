package repository

type User struct {
  Id    string `json:"id,notnull" pg:",pk"`
  Name  string `json:"name,notnull" pg:",notnull"`
	Phone string `json:"phone,notnull" pg:",notnull"`
	Role  string `json:"role,notnull,omitempty" pg:",notnull"`
  Class string `json:"class,omitempty" pg:"-"`
}
