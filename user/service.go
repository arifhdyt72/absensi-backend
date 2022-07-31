package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Insert(input UserInput) (User, error)
	LoginService(input LoginInput) (User, error)
	GetAllUsers(input DatatableInput) ([]User, error)
	GetCountDataUser(input DatatableInput) (int, error)
	DeleteUser(input InputId) (User, error)
	UpdateUser(input UserInput) (User, error)
	GetUserServiceByID(userID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Insert(input UserInput) (User, error) {
	var data User
	data.Username = input.Username
	data.Email = input.Email
	data.Role = input.Role
	data.Fullname = input.Fullname
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return data, err
	}
	data.Password = string(passwordHash)
	data.CreatedAt = time.Now()

	result, err := s.repository.Create(data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) LoginService(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindUserByUsernameOrEmail(username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Empty Username / Email ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, errors.New("Invalid Password")
	}

	return user, nil
}

func (s *service) GetUserServiceByID(userID int) (User, error) {
	result, err := s.repository.FindUserByID(userID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) GetAllUsers(input DatatableInput) ([]User, error) {
	result, err := s.repository.FindAllUsers(input)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) GetCountDataUser(input DatatableInput) (int, error) {
	result, err := s.repository.FindCountDataUser(input)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) DeleteUser(input InputId) (User, error) {
	var user User
	user.ID = input.ID
	result, err := s.repository.Remove(user)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) UpdateUser(input UserInput) (User, error) {
	var data User
	data.ID = input.ID
	data.Username = input.Username
	data.Email = input.Email
	data.Role = input.Role
	data.Fullname = input.Fullname
	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return data, err
		}
		data.Password = string(passwordHash)
	}

	data.UpdatedAt = time.Now()

	result, err := s.repository.Save(data)
	if err != nil {
		return result, err
	}

	return result, nil
}
