package repository

import (
	"time"
)

type Attendance struct {
	tableName struct{}  `pg:"attendance"`
	Id        string    `json:"id"`
	Date      time.Time `json:"date"`
	In        time.Time `json:"in"`
	Out       time.Time `json:"out"`
}

type AttendanceRecordsId struct {
	Id      string       `json:"id"`
	Records []Attendance `json:"records"`
}
