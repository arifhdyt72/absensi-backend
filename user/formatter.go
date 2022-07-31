package user

import "time"

type UserFormat struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUser(data User) UserFormat {
	var result UserFormat
	result.ID = data.ID
	result.Username = data.Username
	result.Email = data.Email
	result.Fullname = data.Fullname
	result.Role = data.Role
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt

	return result
}

func FormatUsers(data []User) []UserFormat {
	var result []UserFormat

	for _, user := range data {
		resultData := FormatUser(user)
		result = append(result, resultData)
	}

	return result
}
