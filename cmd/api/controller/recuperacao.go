package controller

import "github.com/gin-gonic/gin"

func Recuperacao(ctx *gin.Context) {
	// aqui vamos fazer a recuperação de senha, bom podemos fazer da seguuinte forma
	// terá que ser duas paginas para isso, primeiro uma onde ele faz o input do email dele
	// segundo fazemos um disparo de email para ele, para outra pagina onde ele pode colocar a nova senha dele
	// assim fazendo uma rota de edit e mudando a senha do banco de dados
}
