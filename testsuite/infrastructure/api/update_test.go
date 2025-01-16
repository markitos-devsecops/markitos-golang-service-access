package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestUserUpdateHandler_Success(t *testing.T) {
	userRandomId := domain.UUIDv4()
	user := &domain.User{Id: userRandomId, Name: domain.RandomPersonName()}
	userRepository.Create(user)

	updatedName := domain.RandomPersonName()
	requestBody, _ := json.Marshal(map[string]string{
		"name": updatedName,
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/v1/users/"+userRandomId, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response domain.User
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, userRandomId, response.Id)
	assert.Equal(t, updatedName, response.Name)
}

func TestUserUpdateHandler_InvalidID(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{
		"name": "Updated User " + domain.RandomPersonName(),
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/v1/users/invalid-id", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestUserUpdateHandler_MissingName(t *testing.T) {
	user := &domain.User{Id: domain.UUIDv4(), Name: domain.RandomPersonName()}
	userRepository.Create(user)

	requestBody, _ := json.Marshal(map[string]string{})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/v1/users/"+user.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
