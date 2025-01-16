package api_test

import (
	"markitos-golang-service-access/infrastructure/api"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/domain/dependencies"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var userApiServer *api.Server
var userRepository dependencies.UserRepository
var userHasher dependencies.Hasher
var userTokener dependencies.Tokener

func TestMain(m *testing.M) {
	userRepository = domain.NewUserInMemoryRepository()
	userHasher = NewMockSpyUserHasher()
	userTokener = NewMockSpyUserTokener()
	userApiServer = setupTestServer()

	os.Exit(m.Run())
}

func setupTestServer() *api.Server {
	gin.SetMode(gin.TestMode)

	return api.NewServer(":8080", userRepository, userTokener, userHasher)
}
