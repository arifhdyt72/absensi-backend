package attendance

import "time"

type Attendance struct {
	ID             int `gorm:"primary_key:auto_increment"`
	UserId         int
	TimeIn         time.Time
	TimeOut        time.Time
	DateAttendance time.Time
	File           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (Attendance) TableName() string {
	return "attendance"
}

type JoinAttendanceUser struct {
	Username       string
	Email          string
	Fullname       string
	Role           string
	ID             int `gorm:"primary_key:auto_increment"`
	TimeIn         time.Time
	TimeOut        time.Time
	DateAttendance time.Time
	File           string
}
