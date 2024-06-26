package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Saber se o servidor tá on
// @Tags categories
// @Accept json
// @Produce json
// @Failure 400 {object} map[string]interface{} "Erro a entrar no servidor"
// @Router /healthy [get]
func Healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mensage servidor on": "Bomm dia diretoria de gestão de pessoas",
	})
}
