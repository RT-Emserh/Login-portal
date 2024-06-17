package routes

import (
	"github.com/gin-gonic/gin"
	"microservicos.com/cmd/api/controller"
	"microservicos.com/internal/repositories"
	limitbucket "microservicos.com/pkg/middlewares/limitBucket"
)

func CategoryRoutes(r *gin.Engine) {
	r.Use(limitbucket.RateLimiter())
	Categoryroutes := r.Group("/categories")
	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	Categoryroutes.POST("/", func(ctx *gin.Context) {
		controller.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	r.GET("/", controller.ServeIndex)
	r.GET("/d", controller.ServeIndex2)
	r.POST("/login", controller.Login)
	r.GET("/healthy", controller.Healthy)
}

func Healthy(r *gin.Engine) {
}
