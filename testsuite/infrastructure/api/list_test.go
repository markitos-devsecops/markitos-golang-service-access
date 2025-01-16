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

	name1 := domain.RandomPersonName()
	name2 := domain.RandomPersonName()
	user1 := &domain.User{Id: domain.UUIDv4(), Name: name1, Email: domain.RandomEmail(), Password: domain.RandomPassword(10)}
	user2 := &domain.User{Id: domain.UUIDv4(), Name: name2, Email: domain.RandomEmail(), Password: domain.RandomPassword(10)}
	repo.Create(user1)
	repo.Create(user2)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/users/all", nil)

	server.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []*domain.User
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(response), 2)
}
