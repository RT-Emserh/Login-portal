package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservicos.com/cmd/config"
	entites "microservicos.com/internal/entities"
	"microservicos.com/pkg/util"
)

// Login realiza a autenticação do usuário
// @Summary Realiza a autenticação do usuário
// @Description Autentica um usuário com email e senha e retorna um token JWT
// @Tags authentication
// @Accept json
// @Produce json
// @Param credentials body entites.Login true "Credenciais do usuário"
// @Success 200 {object} map[string]interface{} "Token JWT gerado com sucesso"
// @Failure 400 {object} map[string]interface{} "Erro ao realizar login"
// @Failure 401 {object} map[string]interface{} "Credenciais inválidas"
// @Failure 500 {object} map[string]interface{} "Erro no servidor"
// @Router /login [post]
func Login(ctx *gin.Context) {
	var p entites.Login
	var user entites.Category

	db := config.GetDatabase()

	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	// tentei transforma em um middlewares mas não conseguir, tentar novamente pelo turno da noite
	// Check if the user with the provided CPF exists in the database
	if err := db.Where("email = ?", p.Email).First(&user).Error; err != nil {
		fmt.Printf("nome do email %s", p.Email)
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(400, gin.H{
				"error": "user not found",
			})
		} else {
			ctx.JSON(500, gin.H{
				"error": "database error: " + err.Error(),
			})
		}
		return
	}
	fmt.Print(user.Password)
	// Compare the provided password with the stored password hash
	if user.Password != util.Sha256Encoder(p.Password) {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Generate JWT token
	token, err := util.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to generate token: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
