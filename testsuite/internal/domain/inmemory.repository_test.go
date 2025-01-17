package domain_test

import (
	"markitos-golang-service-access/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupInMemoryRepo() *domain.UserInMemoryRepository {
	return domain.NewUserInMemoryRepository()
}

func TestCreate(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Name: "Test User"}
	err := repo.Create(user)
	assert.NoError(t, err)

	result, err := repo.One(&user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, result.Name)
}

func TestDelete(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Name: "Test User"}
	repo.Create(user)

	err := repo.Delete(&user.Id)
	assert.NoError(t, err)

	_, err = repo.One(&user.Id)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Name: "Test User"}
	err := repo.Create(user)
	assert.NoError(t, err)

	updatedName := "Updated User"
	user.Name = updatedName
	err = repo.Update(user)
	assert.NoError(t, err)

	result, err := repo.One(&user.Id)
	assert.NoError(t, err)
	assert.Equal(t, updatedName, result.Name)
}

func TestOne(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Name: "Test User"}
	err := repo.Create(user)
	assert.NoError(t, err)

	result, err := repo.One(&user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, result.Name)
}
