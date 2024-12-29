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

	user := &domain.User{Id: domain.UUIDv4(), Username: domain.RandomEmail()}
	err := repo.Create(user)
	assert.NoError(t, err)

	result, err := repo.One(&user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, result.Username)
}

func TestDelete(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Username: domain.RandomEmail()}
	repo.Create(user)

	err := repo.Delete(&user.Id)
	assert.NoError(t, err)

	_, err = repo.One(&user.Id)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Username: domain.RandomEmail()}
	err := repo.Create(user)
	assert.NoError(t, err)

	updatedMessage := "Updated User"
	user.Username = updatedMessage
	err = repo.Update(user)
	assert.NoError(t, err)

	result, err := repo.One(&user.Id)
	assert.NoError(t, err)
	assert.Equal(t, updatedMessage, result.Username)
}

func TestOne(t *testing.T) {
	repo := setupInMemoryRepo()

	user := &domain.User{Id: domain.UUIDv4(), Username: domain.RandomEmail()}
	err := repo.Create(user)
	assert.NoError(t, err)

	result, err := repo.One(&user.Id)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, result.Username)
}

func TestList(t *testing.T) {
	repo := setupInMemoryRepo()

	email1 := domain.RandomEmail()
	user1 := &domain.User{Id: domain.UUIDv4(), Username: email1}
	email2 := domain.RandomEmail()
	user2 := &domain.User{Id: domain.UUIDv4(), Username: email2}
	err := repo.Create(user1)
	assert.NoError(t, err)
	err = repo.Create(user2)
	assert.NoError(t, err)

	results, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestSearchAndPaginate(t *testing.T) {
	repo := setupInMemoryRepo()

	for i := 0; i < 25; i++ {
		email := "test_search_and_paginate_" + domain.RandomEmail()
		user := &domain.User{Id: domain.UUIDv4(), Username: email}
		err := repo.Create(user)
		assert.NoError(t, err)
	}

	results, err := repo.SearchAndPaginate("test_search_and_paginate_", 2, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 10)
}
