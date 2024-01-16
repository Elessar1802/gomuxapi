package services

import (
	"time"

	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

func PunchIn(db *pg.DB, id string) (interface{}, *err.Error) {
  at := repo.Attendance{Id: id}
  _, er := db.Model(&at).Insert()
  if er != nil {
    return nil, &err.Error{Code: 404, Message: er.Error()}
  }
	return at, nil
}

func PunchOut(db *pg.DB, id string) (interface{}, *err.Error) {
  at := repo.Attendance{Id: id, Date: time.Now(), Out: time.Now()}
  _, er := db.Model(&at).Set("out = ?out").Where("id = ?id").Where("date = ?date").Where("out = null").Update()
  if er != nil {
    // TODO: handle errors in the middleware?
    // FIXME: what about the error codes
    return nil, &err.Error{Code: 404, Message: er.Error()}
  }
	return nil, nil
}

func AttendanceUserId(db *pg.DB, id string, from string, to string) (interface{}, *err.Error) {
  at := []repo.Attendance{}
  er := db.Model(&at).Where("id = ?", id).Where(`"date" >= ?`, from).Where(`"date" <= ?`, to).Select()
  if er != nil {
    return nil, &err.Error{Code: 404, Message: er.Error()}
  }
	return at, nil
}

func AttendanceClassId(db *pg.DB, id string, from string, to string) (interface{}, *err.Error) {
  ids := []string{}
  attendances := []repo.AttendanceRecordsId{}
  // fetch all the student ids in the class
  er := db.Model().Table(id).Select(&ids)
  // for each student id fetch the corresponding attendance records
  for _, i := range ids {
    at, er := AttendanceUserId(db, i, from, to)
    if er != nil {
      return nil, er
    }
    // it is sure to be of the designated type
    t, _ := at.([]repo.Attendance) // type conversion
    attendances = append(attendances, repo.AttendanceRecordsId{Id: i, Records: t})
  }
  if er != nil {
    return nil, &err.Error{Code: 404, Message: er.Error()}
  }
  // attendances is an array of array of attendance records
	return attendances, nil
}
