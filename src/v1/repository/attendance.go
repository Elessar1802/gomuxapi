package repository

import (
	"time"
)

type Attendance struct {
	tableName struct{}  `pg:"attendance"`
	Id        int    `json:"id" pg:",nopk,notnull"` // need to speciy nopk for Fields named Id or Uuid since go-pg is smart
	Date      time.Time `json:"date" pg:"type:date,notnull,default:now()"`
	In        time.Time `json:"in" pg:"type:time,notnull,default:now()"`
	Out       time.Time `json:"out" pg:"type:time"`
}

type AttendanceRecordByDate struct {
  // When someone queries attendance records we will group by id, date & return the following
	Id   int `json:"id"`
	Date string `json:"date"`
	Min  string `json:"first_in"`
	Max  string `json:"last_out"`
	Sum  string `json:"duration"`
}
