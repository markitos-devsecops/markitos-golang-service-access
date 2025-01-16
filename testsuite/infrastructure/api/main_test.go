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

func TestMain(m *testing.M) {
	userApiServer = setupTestServer()
	userRepository = userApiServer.Repository()

	os.Exit(m.Run())
}

func setupTestServer() *api.Server {
	gin.SetMode(gin.TestMode)
	repo := domain.NewUserInMemoryRepository()
	return api.NewServer(":8080", repo)
}
