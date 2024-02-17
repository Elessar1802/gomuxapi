package services

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

const TIMEFORMAT = "2006/01/02"
const MAX_ALLOWED_DAYS = 31
const MAX_ALLOWED_TIME_DIFFERENCE = MAX_ALLOWED_DAYS*24 // 31 days in hours

func PunchIn(db *pg.DB, id string) (encoder.Response) {
  _int_id, _ := strconv.Atoi(id)
  at := repo.Attendance{Id: _int_id, Date: time.Now(), In: time.Now()}

  if p, e := db.Model(&at).Where("id = ?id").Where("date = ?date").Where("out is null").Exists(); p {
    // the user hasn't punched out yet
    if e != nil {
      return err.BadRequestResponse(e.Error())
    }
    return err.BadRequestResponse("User is already punched in")
  }

	_, er := db.Model(&at).Insert()
	if er != nil {
    return err.BadRequestResponse(er.Error())
	}

  return encoder.Response{Code: http.StatusCreated}
}

func PunchOut(db *pg.DB, id string) (encoder.Response) {
  _int_id, _ := strconv.Atoi(id)
	at := repo.Attendance{Id: _int_id, Date: time.Now(), Out: time.Now()}

  if p, _ := db.Model(&at).Where("id = ?id").Where("date = ?date").Where("out is null").Exists(); !p {
    // the user hasn't punched out yet
    return err.BadRequestResponse("User isn't punched in")
  }

	_, er := db.Model(&at).Column("out").Where("id = ?id").Where("date = ?date").Where("out is null").Update()
	if er != nil {
    // most probably user is trying to punch out without punching in
    return err.BadRequestResponse(er.Error())
	}

  // this is from a put request
  return encoder.Response{Code: http.StatusOK}
}

func parseDate(date string) (*time.Time) {
  _date, _e := time.Parse(TIMEFORMAT, date)
  if _e != nil {
    return nil
  }
  return &_date
}

func AttendanceUserId(db *pg.DB, id string, from string, to string) (encoder.Response) {
  at := []repo.Attendance{}
  result := []repo.AttendanceRecordByDate{}
  var _from, _to *time.Time
  if _from = parseDate(from); _from == nil {
    return err.BadRequestResponse("Ill-formed start date")
  }
  if _to = parseDate(to); _to == nil {
    return err.BadRequestResponse("Ill-formed end date")
  }
  dif := _to.Sub(*_from).Hours()
  if (dif < 0) {
    return err.BadRequestResponse("End date can't be after start date")
  } else if (dif > MAX_ALLOWED_TIME_DIFFERENCE) {
    return err.BadRequestResponse(fmt.Sprintf("Specified date range exceeds the limit of %d days", MAX_ALLOWED_DAYS))
  }
	er := db.Model(&at).ColumnExpr("id, date, min(\"in\"), max(out), sum(out - \"in\")").Where("id = ?", id).Where(`"date" >= ?`, from).Where(`"date" <= ?`, to).Group("id", "date").Select(&result)
	if er != nil {
    return err.InternalServerErrorResponse(er.Error())
	}
  return encoder.Response{Code: http.StatusOK, Payload: result}
}

func AttendanceClassId(db *pg.DB, id string, from string, to string) (encoder.Response) {
  attendances := []repo.AttendanceRecordByDate{}
  // Using joins
  _er := GetClass(db, id)
  if _er.Error != nil {
    // propagate the error
    return _er
  }
  var _from, _to *time.Time
  if _from = parseDate(from); _from == nil {
    return err.BadRequestResponse("Ill-formed start date")
  }
  if _to = parseDate(to); _to == nil {
    return err.BadRequestResponse("Ill-formed end date")
  }
  dif := _to.Sub(*_from).Hours()
  if (dif < 0) {
    return err.BadRequestResponse("End date can't be after start date")
  } else if (dif > MAX_ALLOWED_TIME_DIFFERENCE) {
    return err.BadRequestResponse(fmt.Sprintf("Specified date range exceeds the limit of %d days", MAX_ALLOWED_DAYS))
  }
	er := db.Model(&repo.Attendance{}).
		ColumnExpr("attendance.id, attendance.date, min(attendance.\"in\"), max(attendance.out), sum(attendance.out - attendance.\"in\")").
		Join("INNER JOIN students on students.id = attendance.id and students.class = ?", id).
		Where("date >= ?", from).Where("date <= ?", to).Group("attendance.id", "attendance.date").
		Select(&attendances)
	if er != nil {
    return err.InternalServerErrorResponse(er.Error())
	}
	// attendances is an array of array of attendance records
  return encoder.Response{Code: http.StatusOK, Payload: attendances}
}
