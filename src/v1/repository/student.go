package repository

type Student struct {
  Id string `pg:",notnull" json:"id,notnull"`
  Class string `pg:",notnull" json:"class,notnull"`
}
