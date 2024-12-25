package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestUserSearchHandler_Success(t *testing.T) {
	for i := 0; i < 15; i++ {
		message := "Test User " + domain.RandomString(5)
		user := &domain.User{Id: domain.UUIDv4(), Message: message}
		userRepository.Create(user)
	}

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users?search=Test&page=1&size=10", nil)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []*domain.User
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 10)
}

func TestUserSearchHandler_InvalidPageNumber(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users?search=Test&page=invalid&size=10", nil)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestUserSearchHandler_InvalidPageSize(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users?search=Test&page=1&size=invalid", nil)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestUserSearchHandler_EmptyPageNumberItsEqualsToDefaultWithoutErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users?search=Test&page=&size=1", nil)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUserSearchHandler_EmptyPageSizeItsEqualsToDefaultWithoutErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users?search=Test&page=1&size=", nil)

	userApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
