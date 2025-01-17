package services_test

import (
	"markitos-golang-service-access/internal/domain/dependencies"
	"markitos-golang-service-access/internal/services"
	"os"
	"testing"
)

const VALID_UUIDV4 = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
const VALID_NAME = "any valid name"
const VALID_EMAIL = "email@email.com"

var userMockSpyRepository dependencies.UserRepository
var userMockSpyHasher dependencies.Hasher
var userMockSpyTokener dependencies.Tokener
var userCreateService services.UserCreateService
var userOneService services.UserOneService
var userUpdateService services.UserUpdateService
var userLoginService services.UserLoginService

func TestMain(m *testing.M) {
	userMockSpyRepository = NewMockSpyUserRepository()
	userMockSpyHasher = NewMockSpyUserHasher()
	userMockSpyTokener = NewMockSpyUserTokener()

	userCreateService = services.NewUserCreateService(userMockSpyRepository, userMockSpyHasher)
	userOneService = services.NewUserOneService(userMockSpyRepository)
	userUpdateService = services.NewUserUpdateService(userMockSpyRepository)
	userLoginService = services.NewUserLoginService(userMockSpyRepository, userMockSpyHasher, userMockSpyTokener)

	os.Exit(m.Run())
}
