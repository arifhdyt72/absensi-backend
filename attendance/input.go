package attendance

import "time"

type InputAttendance struct {
	ID      int       `json:"id"`
	UserId  int       `json:"user_id" binding:"required"`
	TimeIn  time.Time `json:"time_in"`
	TimeOut time.Time `json:"time_out"`
}

type AttendanceIn struct {
	TimeIn string `form:"time_in" binding:"required"`
}

type AttendanceOut struct {
	ID      int    `json:"id" binding:"required"`
	TimeOut string `json:"time_out" binding:"required"`
}

type DatatableInput struct {
	Filters   string `json:"filters"`
	Page      int    `json:"page"`
	First     int    `json:"first"`
	Rows      int    `json:"rows"`
	PageCount int    `json:"pageCount"`
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
}
