package controller

import (
	"github.com/gin-gonic/gin"
)

// Handler para servir o arquivo HTML
func ServeIndex(ctx *gin.Context) {
	// Define o caminho do arquivo HTML

	// Serve o arquivo HTML usando ctx.File
	ctx.File("index.html")
}

func ServeIndex2(ctx *gin.Context) {
	// Define o caminho do arquivo HTML

	// Serve o arquivo HTML usando ctx.File
	ctx.File("index2.html")
}
