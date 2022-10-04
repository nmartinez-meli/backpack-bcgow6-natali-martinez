package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"mensaje": "Hola Natali",
		})
	})

	groupUsers := router.Group("/users")
	{
		groupUsers.GET("/filter", GetFilterUsers)
		groupUsers.GET("", GetAllUsers)
		groupUsers.GET("/:id", GetUser)
	}

	router.Run()
}

func GetAllUsers(ctx *gin.Context) {
	file, err := os.ReadFile("./ejerciciosTM/users.json")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	if err = json.Unmarshal(file, &users); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, users)
}
func GetUser(ctx *gin.Context) {
	file, err := os.ReadFile("./ejerciciosTM/users.json")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	if err = json.Unmarshal(file, &users); err != nil {
		log.Fatal(err)
	}
	userID := ctx.Param("id")
	userIndex := -1
	for index, user := range users {
		if fmt.Sprintf("%d", user.ID) == userID {
			userIndex = index
		}
	}
	if userIndex < 0 {
		ctx.JSON(404, gin.H{
			"mensaje": "usuario no encontrado",
		})
	} else {

		ctx.JSON(200, users[userIndex])
	}
}
func GetFilterUsers(ctx *gin.Context) {
	file, err := os.ReadFile("./ejerciciosTM/users.json")
	if err != nil {
		log.Fatal(err)
	}
	var users, filterUsers []User
	if err = json.Unmarshal(file, &users); err != nil {
		log.Fatal(err)
	}

	if filterName, isFilterName := ctx.GetQuery("nombre"); isFilterName {
		for _, user := range users {
			if filterName == user.Nombre {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}
	if filterActivo, isFilterActivo := ctx.GetQuery("activo"); isFilterActivo {

		for _, user := range users {
			if filterActivo == fmt.Sprintf("%t", user.Activo) {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}
	if filterApellido, isFilterApellido := ctx.GetQuery("apellido"); isFilterApellido {
		for _, user := range users {
			if filterApellido == user.Apellido {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}
	if filterEmail, isFilterEmail := ctx.GetQuery("email"); isFilterEmail {
		for _, user := range users {
			if filterEmail == user.Email {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}
	if filterFecha, isFilterFecha := ctx.GetQuery("fechaCreacion"); isFilterFecha {
		for _, user := range users {
			if filterFecha == user.FechaCreacion {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}
	if filterEdad, isFilterEdad := ctx.GetQuery("edad"); isFilterEdad {
		for _, user := range users {
			if filterEdad == fmt.Sprintf("%d", user.Edad) {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}
	if filterID, isFilterID := ctx.GetQuery("id"); isFilterID {
		for _, user := range users {
			if filterID == fmt.Sprintf("%d", user.ID) {
				filterUsers = append(filterUsers, user)
			}
		}
		users = filterUsers
		filterUsers = []User{}
	}

	ctx.JSON(200, users)
}
