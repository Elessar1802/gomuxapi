package students

type student struct {
	Id      string `json:"id,notnull"`
	Name    string `json:"name,notnull"`
	Phone   string `json:"phone,notnull"`
	Class   int    `json:"class,notnull"`
	Section string `json:"section,notnull"`
}
