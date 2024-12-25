package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-golang-service-access/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestUserCreateHandler_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserCreateRequest{
		Message: "Test User",
	})
	request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestUserCreateHandler_MissingMessage(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.UserCreateRequest{})
	request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
