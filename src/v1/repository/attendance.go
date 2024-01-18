package repository

import (
	"time"
)

type Attendance struct {
	tableName struct{}  `pg:"attendance"`
  Id        string    `json:"id" pg:",notnull"`
  Date      time.Time `json:"date" pg:"type:date,notnull,default:now()"`
  In        time.Time `json:"in" pg:",notnull,default:now()"`
	Out       time.Time `json:"out"`
}

type AttendanceRecordsId struct {
	Id      string       `json:"id"`
	Records []Attendance `json:"records"`
}
