package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/internal/users"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/pkg/web"
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

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

// @Summary Show a list of users
// @Tags Users
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response "List users"
// @Failure 401 {object} web.Response "Unauthorized"
// @Router /users [GET]
func (c *User) GetAll(ctx *gin.Context) {
	tokenReq := ctx.GetHeader("token")
	if tokenReq != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ""))
		return
	}

	users, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, users, ""))
}

// @Summary  Show a user
// @Tags Users
// @Produce json
// @Param    id     	path 	  int 	        	true   	"Id user"
// @Param    token  	header    string        	true  	"token"
// @Success  200    	{object}  web.Response  			"user details"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  404    	{object}  web.Response  			"user not found"
// @Router   /users/{id} [GET]
func (c *User) GetUser(ctx *gin.Context) {
	tokenReq := ctx.GetHeader("token")
	if tokenReq != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ""))
		return
	}
	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	user, err := c.service.GetUser(int64(userID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
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

// @Summary  Store users
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    users  body      request  true  "User to Store"
// @Success  200    	{object}  web.Response  			"user create"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  400    	{object}  web.Response  			"Bad Request"
// @Router   /users [POST]
func (c *User) CreateUser(ctx *gin.Context) {
	tokenReq := ctx.GetHeader("token")
	if tokenReq != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ""))
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	user, err := c.service.CreateUser(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
}

// @Summary  Update user
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param    id       path      int             true   "User id"
// @Param    token    header    string          false  "Token"
// @Param    user  body      requestUpdateDTO  true   "User to update"
// @Success  200    	{object}  web.Response  			"user update"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  400    	{object}  web.Response  			"Bad Request"
// @Router   /users/{id} [PUT]
func (c *User) Update(ctx *gin.Context) {
	tokenReq := ctx.GetHeader("token")
	if tokenReq != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ""))
		return
	}
	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	var req requestUpdateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	user, err := c.service.Update(req.Nombre, req.Apellido, req.Email, req.FechaCreacion, int64(userID), req.Id, req.Edad, req.Altura, req.Activo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
}

// @Summary  Delete user
// @Tags     Users
// @Param    id     path      int     true  "User id"
// @Param    token  header    string  true  "Token"
// @Success  204
// @Success  200    	{object}  web.Response  			"user delete"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Router   /users/{id} [DELETE]
func (c *User) DeleteUser(ctx *gin.Context) {
	tokenReq := ctx.GetHeader("token")
	if tokenReq != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ""))
		return
	}
	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	if err := c.service.DeleteUser(int64(userID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("el usuario con id = %s se elimino correctamente", paramID), ""))

}

// @Summary      Update name lastname
// @Tags         Users
// @Accept       json
// @Produce      json
// @Description  This endpoint update field name and lastname from a user
// @Param        token  header    string               true  "Token header"
// @Param        id     path      int                  true  "User id"
// @Param        name&lastname  body      requestUpdateFieldsDTO  true  "User name and lastname"
// @Success  200    	{object}  web.Response  			"user update"
// @Failure  401    	{object}  web.Response  			"Unauthorized"
// @Failure  500    	{object}  web.Response  			"Internal server error "
// @Failure  400    	{object}  web.Response  			"Bad Request"
// @Router       /users/{id} [PATCH]
func (c *User) UpdateField(ctx *gin.Context) {
	tokenReq := ctx.GetHeader("token")
	if tokenReq != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ""))
		return
	}
	paramID := ctx.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	var req requestUpdateFieldsDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	user, err := c.service.UpdateField(req.Nombre, req.Apellido, int64(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
}
