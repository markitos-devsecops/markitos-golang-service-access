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
var userRegisterService services.UserRegisterService
var userMeService services.UserMeService
var userUpdateMeService services.UserUpdateMeService
var userLoginService services.UserLoginService

func TestMain(m *testing.M) {
	userMockSpyRepository = NewMockSpyUserRepository()
	userMockSpyHasher = NewMockSpyUserHasher()
	userMockSpyTokener = NewMockSpyUserTokener()

	userRegisterService = services.NewUserRegisterService(userMockSpyRepository, userMockSpyHasher)
	userMeService = services.NewUserMeService(userMockSpyRepository)
	userUpdateMeService = services.NewUserUpdateMeService(userMockSpyRepository)
	userLoginService = services.NewUserLoginService(userMockSpyRepository, userMockSpyHasher, userMockSpyTokener)

	os.Exit(m.Run())
}
