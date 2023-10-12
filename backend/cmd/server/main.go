package main

import (
	"backend/services"
	"log"

	_ "backend/cmd/server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	app := gin.Default()
	app.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
		},
	))

	services, err := services.NewService()
	if err != nil {
		log.Fatal("Failed to init service", err)
	}

	app.GET("/product", services.GetListProduct)
	app.GET("/product/:id", services.GetSingleProduct)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run("127.0.0.1:8080")
}
