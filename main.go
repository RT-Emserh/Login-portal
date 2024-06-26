package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"microservicos.com/cmd/config"
	"microservicos.com/cmd/routes"
	_ "microservicos.com/docs"
	entites "microservicos.com/internal/entities"
)

// @title Login Portal do colaborador
// @version 1.0
// @description essa é a documentação do login do portal do colaborador.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.basic BasicAuth
func main() {
	config.Connect()
	config.DB.AutoMigrate(entites.Category{})
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.Healthy(r)
	routes.CategoryRoutes(r)

	r.Run(":5000")
}
