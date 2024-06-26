package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"microservicos.com/internal/repositories"
	usecases "microservicos.com/internal/use-cases"
	"microservicos.com/pkg/processors"
)

// aqui vamos utilizar como dto, por isso está em minusculo
type createCategoryInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Cpf      string `json:"cpf" binding:"required"`
	Password string `json:"password" binding:"required"`
	Telefone string `json:"telefone" binding:"required"`
}

// CreateCategory cria uma nova categoria
// @Summary Cria uma nova categoria
// @Description Cria uma nova categoria com os detalhes fornecidos
// @Tags categories
// @Accept json
// @Produce json
// @Param category body createCategoryInput true "Detalhes da nova categoria"
// @Success 200 {object} map[string]interface{} "Cadastrado com sucesso"
// @Failure 400 {object} map[string]interface{} "Erro ao cadastrar"
// @Router /categories [post]
func CreateCategory(ctx *gin.Context, repository repositories.IRepositories) {
	var body createCategoryInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message error": err.Error(),
		})
		return
	}
	useCase := usecases.NewcreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name, body.Email, body.Telefone, body.Cpf, body.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	go processors.Processors()
	processors.EnqueueEmailJob(body.Email, body.Name, body.Cpf)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cadastrado com sucesso, o email cadastrado no portal será o email a qual será enviado o contracheque e outros recebibos.",
	})
}
