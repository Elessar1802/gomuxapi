package repository

type Credential struct {
  Id       int `json:"id,notnull" pg:"type:integer,nopk,notnull"` // this will be a foreign key
	Password string `json:"password,notnull" pg:",notnull"`
	Token    string `json:"-"`
	Role     string `json:"role,notnull" pg:",notnull"`
}
