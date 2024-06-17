package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"microservicos.com/cmd/routes"
	entities "microservicos.com/internal/entities"
)

func TestCreateCategory(t *testing.T) {
	r := gin.Default()
	routes.CategoryRoutes(r)

	category := entities.Category{
		Name:      "Demo name",
		Email:     "js0454261@gmail.com",
		Telefone:  "98984369945",
		Cpf:       "6202073535",
		Password:  "123456",
		CreatedAt: time.Now(),
	}

	jsonValue, _ := json.Marshal(category)

	request, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(jsonValue))

	W := httptest.NewRecorder()
	r.ServeHTTP(W, request)

	// Verifica se o status é 307 (Redirecionamento temporário)
	assert.Equal(t, http.StatusTemporaryRedirect, W.Code)
}
