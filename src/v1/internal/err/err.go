package err

type Error struct {
	Code    int    `json:"code,notnull"`
	Message string `json:"message,notnull"`
}
