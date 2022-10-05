package users

import (
	"fmt"
	"time"
)

type Service interface {
	GetAll() ([]User, error)
	CreateUser(nombre, apellido, email string, edad int64, altura float64, activo bool) (User, error)
	GetUser(id int64) (User, error)
	Update(nombre, apellido, email, fechaCreacion string, paramID, id, edad int64, altura float64, activo bool) (User, error)
	UpdateField(nombre, apellido string, paramID int64) (User, error)
	DeleteUser(paramID int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *service) GetUser(id int64) (User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *service) CreateUser(nombre, apellido, email string, edad int64, altura float64, activo bool) (User, error) {
	y, m, d := time.Now().Date()
	fechaCreacion := fmt.Sprintf("%d-%2d-%4d", d, m, y)
	user, err := s.repository.CreateUser(nombre, apellido, email, fechaCreacion, edad, altura, activo)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (s *service) Update(nombre, apellido, email, fechaCreacion string, paramID, id, edad int64, altura float64, activo bool) (User, error) {
	user, err := s.repository.GetUser(paramID)
	if err != nil {
		return User{}, err
	}
	user, err = s.repository.Update(nombre, apellido, email, fechaCreacion, id, edad, altura, activo, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (s *service) UpdateField(nombre, apellido string, paramID int64) (User, error) {
	user, err := s.repository.GetUser(paramID)
	if err != nil {
		return User{}, err
	}
	user, err = s.repository.UpdateField(nombre, apellido, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (s *service) DeleteUser(paramID int64) error {
	user, err := s.repository.GetUser(paramID)
	if err != nil {
		return err
	}

	if err := s.repository.DeleteUser(user); err != nil {
		return err
	}
	return nil
}
