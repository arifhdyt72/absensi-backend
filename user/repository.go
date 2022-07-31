package user

import "gorm.io/gorm"

type Repository interface {
	Create(data User) (User, error)
	FindUserByUsernameOrEmail(user string) (User, error)
	FindAllUsers(input DatatableInput) ([]User, error)
	FindCountDataUser(input DatatableInput) (int, error)
	Save(data User) (User, error)
	Remove(data User) (User, error)
	FindUserByID(userID int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(data User) (User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindUserByUsernameOrEmail(user string) (User, error) {
	var data User
	err := r.db.Where("username = ? OR email = ?", user, user).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindAllUsers(input DatatableInput) ([]User, error) {
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

	var data []User
	err := r.db.Where("username LIKE ? OR email LIKE ? OR fullname LIKE ? OR role LIKE ?", search, search, search, search).
		Limit(input.Rows).Offset(input.First).Order(input.SortField + " " + input.SortOrder).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindCountDataUser(input DatatableInput) (int, error) {
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

	var data []User
	err := r.db.Where("username LIKE ? OR email LIKE ? OR fullname LIKE ? OR role LIKE ?", search, search, search, search).
		Limit(input.Rows).Offset(input.First).Order(input.SortField + " " + input.SortOrder).Find(&data).Count(&count).Error
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}

func (r *repository) Save(data User) (User, error) {
	if data.Password == "" {
		err := r.db.Model(User{}).Where("id = ?", data.ID).
			Updates(&User{
				Username:  data.Username,
				Email:     data.Email,
				Fullname:  data.Fullname,
				Role:      data.Role,
				UpdatedAt: data.UpdatedAt,
			}).Error
		if err != nil {
			return data, err
		}

		return data, nil
	} else {
		err := r.db.Model(User{}).Where("id = ?", data.ID).
			Updates(&User{
				Username:  data.Username,
				Password:  data.Password,
				Email:     data.Email,
				Fullname:  data.Fullname,
				Role:      data.Role,
				UpdatedAt: data.UpdatedAt,
			}).Error
		if err != nil {
			return data, err
		}

		return data, nil
	}
}

func (r *repository) Remove(data User) (User, error) {
	err := r.db.Delete(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) FindUserByID(userID int) (User, error) {
	var data User
	err := r.db.Where("id = ?", userID).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
