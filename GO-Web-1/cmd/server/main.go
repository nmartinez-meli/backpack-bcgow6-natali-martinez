package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/cmd/server/handler"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/internal/users"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	user := handler.NewProduct(service)

	router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"mensaje": "Hola Natali",
	// 	})
	// })

	groupUsers := router.Group("/users")
	{
		// groupUsers.GET("/filter", user.GetFilterUsers)
		groupUsers.GET("", user.GetAll)
		groupUsers.GET("/:id", user.GetUser)
		groupUsers.POST("", user.CreateUser)
		groupUsers.PUT("/:id", user.Update)
		groupUsers.PATCH("/:id", user.UpdateField)
		groupUsers.DELETE("/:id", user.DeleteUser)
	}

	router.Run()
}
