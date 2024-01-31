package repository

import (
	"time"
)

type Attendance struct {
	tableName struct{}  `pg:"attendance"`
  Id        string    `json:"id" pg:",nopk,notnull"` // need to speciy nopk for Fields named Id or Uuid since go-pg is smart
  Date      time.Time `json:"date" pg:"type:date,notnull,default:now()"`
  In        time.Time `json:"in" pg:",notnull,default:now()"`
	Out       time.Time `json:"out"`
}

type AttendanceRecordsId struct {
	Id      string       `json:"id"`
	Records []Attendance `json:"records"`
}
