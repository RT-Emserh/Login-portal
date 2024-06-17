package test

/*
func TestLogin(t *testing.T) {
	var user entites.Login
	gin.SetMode(gin.TestMode)

	config.DB.Find(user)

	// Estrutura de login para enviar na solicitação
	loginPayload := map[string]string{
		"email":   "jls.silva@discente.ufma.br",
		"passwor": "123456",
	}
	jsonValue, _ := json.Marshal(loginPayload)

	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/login", controller.Login)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response["token"])
}
*/
