package api_test

import (
	"bytes"
	"encoding/json"
	"markitos-golang-service-access/infrastructure/api"
	"markitos-golang-service-access/infrastructure/implementations"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var userApiServer *api.Server
var userRepository dependencies.UserRepository
var userHasher dependencies.Hasher
var userTokener dependencies.Tokener

func TestMain(m *testing.M) {
	userRepository = domain.NewUserInMemoryRepository()
	userHasher = implementations.NewHasherBCrypt()
	userTokener, _ = implementations.NewTokenerPasseto("12345678901234567890123456789012")
	userApiServer = setupTestServer()

	os.Exit(m.Run())
}

func setupTestServer() *api.Server {
	gin.SetMode(gin.TestMode)

	return api.NewServer(":8080", userRepository, userTokener, userHasher)
}

func createUserAndLogin(t *testing.T) (*domain.User, string) {
	temporalyId := domain.UUIDv4()
	user, _ := domain.NewUser(temporalyId, domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
	registerRequestBody, _ := json.Marshal(map[string]string{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/v1/users/register", bytes.NewBuffer(registerRequestBody))
	request.Header.Set("Content-Type", "application/json")
	userApiServer.Router().ServeHTTP(recorder, request)
	require.Equal(t, http.StatusCreated, recorder.Code)
	var registerResponse map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &registerResponse)
	require.NoError(t, err)
	id, ok := registerResponse["id"]
	require.True(t, ok)
	user.Id = id

	loginRequestBody, _ := json.Marshal(map[string]string{
		"email":    user.Email,
		"password": user.Password,
	})

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest(http.MethodPost, "/v1/users/login", bytes.NewBuffer(loginRequestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	var loginResponse map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &loginResponse)
	require.NoError(t, err)
	token, ok := loginResponse["token"]
	require.True(t, ok)

	return user, token
}
