package api_test

import (
	"bytes"
	"encoding/json"
	"markitos-golang-service-access/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUpdateMeHandler_Success(t *testing.T) {
	user, token := createUserAndLogin(t)

	updatedName := domain.RandomPersonName()
	requestBody, _ := json.Marshal(map[string]string{
		"name": updatedName,
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/v1/users/me", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response domain.User
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, response.Id)
	assert.Equal(t, updatedName, response.Name)
	assert.NotEqual(t, user.Name, response.Name)
}

func TestUserUpdateMeHandler_MissingName(t *testing.T) {
	_, token := createUserAndLogin(t)

	requestBody, _ := json.Marshal(map[string]string{})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/v1/users/me", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
