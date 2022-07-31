package user

import "time"

type User struct {
	ID        int `gorm:"primary_key:auto_increment"`
	Username  string
	Password  string
	Email     string
	Fullname  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "master_user"
}
