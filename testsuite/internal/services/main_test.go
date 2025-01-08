package services_test

import (
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/services"
	"os"
	"testing"
)

const VALID_UUIDV4 = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
const VALID_NAME = "any valid name"

var userMockSpyRepository domain.UserRepository
var userCreateService services.UserCreateService
var userOneService services.UserOneService
var userListService services.UserListService
var userUpdateService services.UserUpdateService

func TestMain(m *testing.M) {
	userMockSpyRepository = NewMockSpyUserRepository()
	userCreateService = services.NewUserCreateService(userMockSpyRepository)
	userOneService = services.NewUserOneService(userMockSpyRepository)
	userListService = services.NewUserListService(userMockSpyRepository)
	userUpdateService = services.NewUserUpdateService(userMockSpyRepository)

	os.Exit(m.Run())
}

type MockSpyUserRepository struct {
	LastCreatedUser       *domain.User
	LastCreatedForOneUser *domain.User
	OneCalled             bool
	LastUpdatedUser       *domain.User
}

func NewMockSpyUserRepository() *MockSpyUserRepository {
	return &MockSpyUserRepository{
		LastCreatedUser:       nil,
		LastCreatedForOneUser: nil,
		OneCalled:             false,
		LastUpdatedUser:       nil,
	}
}

func (m *MockSpyUserRepository) Create(user *domain.User) error {
	m.LastCreatedUser = user
	m.LastCreatedForOneUser = user

	return nil
}

func (m *MockSpyUserRepository) CreateHaveBeenCalledWith(user *domain.User) bool {
	var result bool = m.LastCreatedUser.Id == user.Id && m.LastCreatedUser.Name == user.Name

	m.LastCreatedUser = nil

	return result
}

func (m *MockSpyUserRepository) CreateHaveBeenCalledWithName(user *domain.User) bool {
	var result bool = m.LastCreatedUser.Name == user.Name

	m.LastCreatedUser = nil

	return result
}

func (m *MockSpyUserRepository) Delete(id *string) error {
	return nil
}

func (m *MockSpyUserRepository) Update(user *domain.User) error {
	m.LastUpdatedUser = user

	return nil
}

func (m *MockSpyUserRepository) One(id *string) (*domain.User, error) {
	return &domain.User{
		Id:   *id,
		Name: VALID_NAME,
	}, nil
}

func (m *MockSpyUserRepository) SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*domain.User, error) {
	return []*domain.User{
		{
			Id:   VALID_UUIDV4,
			Name: VALID_NAME,
		},
	}, nil
}

func (m *MockSpyUserRepository) OneHaveBeenCalledWith(user *domain.User) bool {
	var result bool = m.LastCreatedForOneUser.Id == user.Id && m.LastCreatedForOneUser.Name == user.Name

	m.LastCreatedForOneUser = nil

	return result
}

func (m *MockSpyUserRepository) OneHaveBeenCalledWithName(user *domain.User) bool {
	var result bool = m.LastCreatedForOneUser.Name == user.Name && m.LastCreatedForOneUser.Id == user.Id

	m.LastCreatedUser = nil

	return result
}

func (m *MockSpyUserRepository) UpdateHaveBeenCalledWithName(user *domain.User) bool {
	var result bool = m.LastUpdatedUser.Name == user.Name && m.LastUpdatedUser.Id == user.Id

	m.LastUpdatedUser = nil

	return result
}

func (m *MockSpyUserRepository) List() ([]*domain.User, error) {
	m.OneCalled = true

	return []*domain.User{}, nil
}
func (m *MockSpyUserRepository) ListHaveBeenCalled() bool {
	var result bool = m.OneCalled

	m.OneCalled = false

	return result
}
