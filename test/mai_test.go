package test

import "github.com/gin-gonic/gin"

func SetUprouter() *gin.Engine {
	router := gin.Default()
	return router
}
