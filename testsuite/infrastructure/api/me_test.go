package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserMeHandler_Success(t *testing.T) {
	user, token := createUserAndLogin(t)
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/users/me", nil)
	request.Header.Set("Authorization", "Bearer "+token)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response domain.User
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, response.Id)
	assert.Equal(t, user.Name, response.Name)
	assert.Equal(t, user.Email, response.Email)
}

func TestUserMeHandler_Unauthorized(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/users/me", nil)
	request.Header.Set("Authorization", "Bearer "+domain.RandomString(32))

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	require.Empty(t, recorder.Header().Get("Authorization"))
}

func TestUserMeHandler_WithoutAuthorization(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/users/me", nil)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	require.Empty(t, recorder.Header().Get("Authorization"))
}
