package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-golang-service-access/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestUserListHandler_Success(t *testing.T) {
	server := setupTestServer()
	repo := server.Repository().(*domain.UserInMemoryRepository)

	user1 := &domain.User{Id: domain.UUIDv4(), Message: "Test User 1"}
	user2 := &domain.User{Id: domain.UUIDv4(), Message: "Test User 2"}
	repo.Create(user1)
	repo.Create(user2)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/users/all", nil)

	server.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []*domain.User
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "Test User 1", response[0].Message)
	assert.Equal(t, "Test User 2", response[1].Message)
}
