package attendance

import (
	"absensi-backend/helper"
	"absensi-backend/user"
	"time"
)

type Service interface {
	InsertAttendance(input InputAttendance) (Attendance, error)
	GetAttendanceNowByUserID(user_id int) (Attendance, error)
	AttendanceInService(user_id int, input AttendanceIn, nameFile string) (Attendance, error)
	AttendanceOutService(input AttendanceOut) (Attendance, error)
	GetAllAttendance(input DatatableInput, user user.User) ([]JoinAttendanceUser, error)
	GetCountDataAttendance(input DatatableInput, user user.User) (int, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InsertAttendance(input InputAttendance) (Attendance, error) {
	var data Attendance
	data.UserId = input.UserId
	data.TimeIn = input.TimeIn
	data.DateAttendance = time.Now()

	result, err := s.repository.Create(data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) GetAttendanceNowByUserID(user_id int) (Attendance, error) {
	result, err := s.repository.FindByNowDateAndUserID(user_id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) AttendanceInService(user_id int, input AttendanceIn, nameFile string) (Attendance, error) {
	var data Attendance
	timeIn, _ := helper.FormatTime(input.TimeIn)

	data.UserId = user_id
	data.TimeIn = timeIn
	data.DateAttendance = time.Now()
	data.File = nameFile
	data.CreatedAt = time.Now()

	result, err := s.repository.Create(data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) AttendanceOutService(input AttendanceOut) (Attendance, error) {
	var data Attendance
	timeOut, _ := helper.FormatTime(input.TimeOut)

	data.ID = input.ID
	data.TimeOut = timeOut
	data.UpdatedAt = time.Now()

	result, err := s.repository.Update(data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) GetAllAttendance(input DatatableInput, user user.User) ([]JoinAttendanceUser, error) {
	result, err := s.repository.FindAllAttendance(input, user)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) GetCountDataAttendance(input DatatableInput, user user.User) (int, error) {
	result, err := s.repository.FindCountDataAttendance(input, user)
	if err != nil {
		return result, err
	}

	return result, nil
}
