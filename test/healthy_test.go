package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"microservicos.com/cmd/api/controller"
)

func TestHomepageHandler(t *testing.T) {
	mockResponse := `{"mensage servidor on":"Bomm dia diretoria de gestão de pessoas"}`
	r := SetUprouter()
	r.GET("/healthy", controller.Healthy)
	req, _ := http.NewRequest("GET", "/healthy", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
