package users

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

type User struct {
	ID            int64   `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int64   `json:"edad"`
	Activo        bool    `json:"activo"`
	Altura        float64 `json:"altura"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var users []User

type Repository interface {
	GetAll() ([]User, error)
	CreateUser(nombre, apellido, email, fechaCreacion string, edad int64, altura float64, activo bool) (User, error)
	GetUser(id int64) (User, error)
	Update(nombre, apellido, email, fechaCreacion string, id, edad int64, altura float64, activo bool, user User) (User, error)
	UpdateField(nombre, apellido string, user User) (User, error)
	DeleteUser(user User) error
	// Update(id int, name, productType string, count int, price float64) (User, error)
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	file, err := os.ReadFile("../../users.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(file, &users); err != nil {
		log.Fatal(err)
	}
	return &repository{}
}

func (r *repository) CreateUser(nombre, apellido, email, fechaCreacion string, edad int64, altura float64, activo bool) (User, error) {
	user := User{-1, nombre, apellido, email, edad, activo, altura, fechaCreacion}
	if id, err := lastID(); err != nil {
		return user, err
	} else {
		user.ID = id + 1
	}
	users = append(users, user)
	if err := write(users); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) Update(nombre, apellido, email, fechaCreacion string, id, edad int64, altura float64, activo bool, user User) (User, error) {
	auxID := user.ID
	user = User{id, nombre, apellido, email, edad, activo, altura, fechaCreacion}
	for index := range users {
		if users[index].ID == auxID {
			users[index] = user
		}
	}
	if err := write(users); err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *repository) UpdateField(nombre, apellido string, user User) (User, error) {
	user.Nombre = nombre
	user.Apellido = apellido

	for index := range users {
		if users[index].ID == user.ID {
			users[index] = user
		}
	}
	if err := write(users); err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) GetUser(id int64) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("usuario no encontrado")

}
func (r *repository) DeleteUser(user User) error {
	var filter []User
	for _, u := range users {
		if u.ID != user.ID {
			filter = append(filter, u)
		}
	}
	users = filter
	if err := write(users); err != nil {
		return err
	}
	return nil

}

func lastID() (int64, error) {
	return users[len(users)-1].ID, nil
}

func write(users []User) error {
	jsonUser, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return os.WriteFile("../../users.json", []byte(strings.ReplaceAll(string(jsonUser), ",", ",\n")), 0644)
}
