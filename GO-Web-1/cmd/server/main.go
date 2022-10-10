package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/cmd/server/handler"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/docs"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/internal/users"
	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/GO-Web-1/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp Go W6 - API
// @version 4.2
// @description This API handle MELI users
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name Natali Martinez
// @contact.url natali.martinez@mercadolibre.com
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	store := store.New(store.FileType, os.Getenv("JSON_STORE_FILE"))
	repo := users.NewRepository(store)
	service := users.NewService(repo)
	user := handler.NewUser(service)

	router := gin.Default()
	// Documentación swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
