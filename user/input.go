package user

import "time"

type UserInput struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password"`
	Email     string    `json:"email" binding:"required"`
	Fullname  string    `json:"fullname" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputId struct {
	ID int `uri:"id" binding:"required"`
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
