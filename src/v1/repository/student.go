package repository

type Student struct {
  Id int `pg:",nopk,notnull" json:"id,notnull"`
  Class string `pg:",nopk,notnull" json:"class,notnull"`
}
