package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/internal/users"
)

type request struct {
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int64   `json:"edad" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	FechaCreacion string  `json:"fechaCreacion"`
}
type requestUpdateDTO struct {
	Id            int64   `json:"id" binding:"required"`
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int64   `json:"edad" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	FechaCreacion string  `json:"fechaCreacion"`
}
type requestUpdateFieldsDTO struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
}

type User struct {
	service users.Service
}

func NewProduct(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetAll(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	user, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, user)
}

func (c *User) GetUser(ctx *gin.Context) {

	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := c.service.GetUser(int64(userID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

// func (c *User) GetFilterUsers(ctx *gin.Context) {

// 	// if filterName, isFilterName := ctx.GetQuery("nombre"); isFilterName {
// 	// 	for _, user := range users {
// 	// 		if filterName == user.Nombre {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }
// 	// if filterActivo, isFilterActivo := ctx.GetQuery("activo"); isFilterActivo {

// 	// 	for _, user := range users {
// 	// 		if filterActivo == fmt.Sprintf("%t", user.Activo) {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }
// 	// if filterApellido, isFilterApellido := ctx.GetQuery("apellido"); isFilterApellido {
// 	// 	for _, user := range users {
// 	// 		if filterApellido == user.Apellido {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }
// 	// if filterEmail, isFilterEmail := ctx.GetQuery("email"); isFilterEmail {
// 	// 	for _, user := range users {
// 	// 		if filterEmail == user.Email {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }
// 	// if filterFecha, isFilterFecha := ctx.GetQuery("fechaCreacion"); isFilterFecha {
// 	// 	for _, user := range users {
// 	// 		if filterFecha == user.FechaCreacion {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }
// 	// if filterEdad, isFilterEdad := ctx.GetQuery("edad"); isFilterEdad {
// 	// 	for _, user := range users {
// 	// 		if filterEdad == fmt.Sprintf("%d", user.Edad) {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }
// 	// if filterID, isFilterID := ctx.GetQuery("id"); isFilterID {
// 	// 	for _, user := range users {
// 	// 		if filterID == fmt.Sprintf("%d", user.ID) {
// 	// 			filterUsers = append(filterUsers, user)
// 	// 		}
// 	// 	}
// 	// 	users = filterUsers
// 	// 	filterUsers = []User{}
// 	// }

// 	// ctx.JSON(200, users)
// 	ctx.JSON(200, gin.H{
// 		"mensaje": "Hola Natali desde filtro",
// 	})
// }

func (c *User) CreateUser(ctx *gin.Context) {

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := c.service.CreateUser(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
func (c *User) Update(ctx *gin.Context) {

	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var req requestUpdateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := c.service.Update(req.Nombre, req.Apellido, req.Email, req.FechaCreacion, int64(userID), req.Id, req.Edad, req.Altura, req.Activo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (c *User) DeleteUser(ctx *gin.Context) {

	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.service.DeleteUser(int64(userID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"mensaje": fmt.Sprintf("el uusuario con id = %s se elimino correctamente", paramID),
	})
}
func (c *User) UpdateField(ctx *gin.Context) {

	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var req requestUpdateFieldsDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := c.service.UpdateField(req.Nombre, req.Apellido, int64(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
