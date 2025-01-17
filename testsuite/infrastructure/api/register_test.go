package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"

	"github.com/stretchr/testify/assert"
)

const USER_REGISTER_ENDPOINT = "/v1/users/register"

func TestUserRegisterHandler_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserRegisterRequest{
		Name:     domain.RandomPersonName(),
		Email:    domain.RandomEmail(),
		Password: domain.RandomPassword(10),
	})
	request, _ := http.NewRequest(http.MethodPost, USER_REGISTER_ENDPOINT, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestUserRegisterHandler_MissingName(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserRegisterRequest{
		Email:    domain.RandomEmail(),
		Password: domain.RandomPassword(10),
	})
	request, _ := http.NewRequest(http.MethodPost, USER_REGISTER_ENDPOINT, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestUserRegisterHandler_MissingEmail(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserRegisterRequest{
		Name:     domain.RandomPersonName(),
		Password: domain.RandomPassword(10),
	})
	request, _ := http.NewRequest(http.MethodPost, USER_REGISTER_ENDPOINT, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestUserRegisterHandler_NotAnEmail(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserRegisterRequest{
		Email:    "not-an-email",
		Name:     domain.RandomPersonName(),
		Password: domain.RandomPassword(10),
	})
	request, _ := http.NewRequest(http.MethodPost, USER_REGISTER_ENDPOINT, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestUserRegisterHandler_MissingPassword(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserRegisterRequest{
		Name:  domain.RandomPersonName(),
		Email: domain.RandomEmail(),
	})
	request, _ := http.NewRequest(http.MethodPost, USER_REGISTER_ENDPOINT, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
