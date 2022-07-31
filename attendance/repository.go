package attendance

import (
	"absensi-backend/user"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(data Attendance) (Attendance, error)
	Update(data Attendance) (Attendance, error)
	FindByNowDateAndUserID(userID int) (Attendance, error)
	FindAttendanceByID(ID int) (Attendance, error)
	FindAllAttendance(input DatatableInput, user user.User) ([]JoinAttendanceUser, error)
	FindCountDataAttendance(input DatatableInput, user user.User) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(data Attendance) (Attendance, error) {
	err := r.db.Model(Attendance{}).
		Create(map[string]interface{}{
			"user_id":         data.UserId,
			"time_in":         data.TimeIn.Format("2006-01-02 15:04:05"),
			"date_attendance": data.DateAttendance,
			"file":            data.File,
			"created_at":      data.CreatedAt,
		}).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data Attendance) (Attendance, error) {
	err := r.db.Model(Attendance{}).Where("ID = ?", data.ID).
		Updates(map[string]interface{}{
			"time_out":   data.TimeOut.Format("2006-01-02 15:04:05"),
			"updated_at": data.UpdatedAt,
		}).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindByNowDateAndUserID(userID int) (Attendance, error) {
	dateNow := time.Now().Format("2006-01-02")
	var data Attendance
	err := r.db.Where("user_id = ? AND date_attendance = ?", userID, dateNow).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindAttendanceByID(ID int) (Attendance, error) {
	var data Attendance
	err := r.db.Where("ID = ?", ID).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindAllAttendance(input DatatableInput, user user.User) ([]JoinAttendanceUser, error) {
	if input.Rows == 0 {
		input.Rows = 10
	}

	if input.SortField == "" || input.SortField == "null" {
		input.SortField = "id"
	}

	if input.SortOrder == "1" {
		input.SortOrder = "ASC"
	} else {
		input.SortOrder = "DESC"
	}

	search := "%" + input.Filters + "%"
	var data []JoinAttendanceUser
	if user.Role == "admin" {
		err := r.db.Table("attendance a").Select(`
			u.username,
			u.email,
			u.fullname,
			u.role,
			a.ID,
			a.time_in,
			a.time_out,
			a.date_attendance,
			a.file
		`).Joins("INNER JOIN master_user u ON a.user_id = u.id").
			Where("u.username LIKE ? OR u.email LIKE ? OR u.fullname LIKE ? OR u.role LIKE ? OR a.file LIKE ? OR a.date_attendance LIKE ?",
				search, search, search, search, search, search).Limit(input.Rows).Offset(input.First).
			Order(input.SortField + " " + input.SortOrder).Find(&data).Error
		if err != nil {
			return data, err
		}

		return data, nil
	} else {
		err := r.db.Table("attendance a").Select(`
		u.username,
		u.email,
		u.fullname,
		u.role,
		a.ID,
		a.time_in,
		a.time_out,
		a.date_attendance,
		a.file
	`).Joins("INNER JOIN master_user u ON a.user_id = u.id").Where("a.user_id = ?", user.ID).
			Where("u.username LIKE ? OR u.email LIKE ? OR u.fullname LIKE ? OR u.role LIKE ? OR a.file LIKE ? OR a.date_attendance LIKE ?",
				search, search, search, search, search, search).Limit(input.Rows).Offset(input.First).
			Order(input.SortField + " " + input.SortOrder).Find(&data).Error
		if err != nil {
			return data, err
		}

		return data, nil
	}
}

func (r *repository) FindCountDataAttendance(input DatatableInput, user user.User) (int, error) {
	if input.Rows == 0 {
		input.Rows = 10
	}

	if input.SortField == "" || input.SortField == "null" {
		input.SortField = "id"
	}

	if input.SortOrder == "1" {
		input.SortOrder = "ASC"
	} else {
		input.SortOrder = "DESC"
	}

	var count int64
	search := "%" + input.Filters + "%"
	var data []JoinAttendanceUser
	if user.Role == "admin" {
		err := r.db.Table("attendance a").Select(`
			u.username,
			u.email,
			u.fullname,
			u.role,
			a.ID,
			a.time_in,
			a.time_out,
			a.date_attendance,
			a.file
		`).Joins("INNER JOIN master_user u ON a.user_id = u.id").
			Where("u.username LIKE ? OR u.email LIKE ? OR u.fullname LIKE ? OR u.role LIKE ? OR a.file LIKE ? OR a.date_attendance LIKE ?",
				search, search, search, search, search, search).Limit(input.Rows).
			Offset(input.First).Order(input.SortField + " " + input.SortOrder).Find(&data).Count(&count).Error
		if err != nil {
			return int(count), err
		}

		return int(count), nil
	} else {
		err := r.db.Table("attendance a").Select(`
		u.username,
		u.email,
		u.fullname,
		u.role,
		a.ID,
		a.time_in,
		a.time_out,
		a.date_attendance,
		a.file
	`).Joins("INNER JOIN master_user u ON a.user_id = u.id").Where("a.user_id = ?", user.ID).
			Where("u.username LIKE ? OR u.email LIKE ? OR u.fullname LIKE ? OR u.role LIKE ? OR a.file LIKE ? OR a.date_attendance LIKE ?",
				search, search, search, search, search, search).Limit(input.Rows).
			Offset(input.First).Order(input.SortField + " " + input.SortOrder).Find(&data).Count(&count).Error
		if err != nil {
			return int(count), err
		}

		return int(count), nil
	}
}
