package database_test

import (
	"log"
	"markitos-golang-service-access/infrastructure/configuration"
	"markitos-golang-service-access/infrastructure/database"
	"markitos-golang-service-access/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserCreate(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), "Hello, World!")
	err := repository.Create(user)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", user.Id).Error
	require.NoError(t, err)
	require.Equal(t, user.Id, result.Id)
	require.Equal(t, user.Message, result.Message)
	require.WithinDuration(t, user.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, user.UpdatedAt, result.UpdatedAt, time.Second)

	db.Delete(&result)
}

func TestSearch(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewUserPostgresRepository(db)

	randomMessage := "Test " + domain.RandomString(10)
	user := &domain.User{Id: domain.UUIDv4(), Message: randomMessage}
	db.Create(user)

	results, err := repo.SearchAndPaginate(randomMessage, 1, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, randomMessage, results[0].Message)

	cleanDB(db)
}
func TestUserDelete(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), "Hello, World!")
	db.Create(user)

	err := repository.Delete(&user.Id)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", user.Id).Error
	require.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUserUpdate(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), "Hello, World!")
	db.Create(user)

	user.Message = "Updated Message"
	err := repository.Update(user)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", user.Id).Error
	require.NoError(t, err)
	require.Equal(t, "Updated Message", result.Message)
}

func TestUserOne(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), "Hello, World!")
	db.Create(user)

	result, err := repository.One(&user.Id)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, user.Id, result.Id)
	require.Equal(t, user.Message, result.Message)
}

func TestUserList(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repository := database.NewUserPostgresRepository(db)

	for i := 0; i < 5; i++ {
		user := &domain.User{Id: domain.UUIDv4(), Message: domain.RandomString(10)}
		db.Create(user)
	}

	results, err := repository.List()
	require.NoError(t, err)
	require.Len(t, results, 5)
}
func TestPagination(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewUserPostgresRepository(db)

	for i := 0; i < 15; i++ {
		user := &domain.User{Id: domain.UUIDv4(), Message: domain.RandomString(10)}
		db.Create(user)
	}

	results, err := repo.SearchAndPaginate("", 2, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 5)

	cleanDB(db)
}

func TestSearchAndPagination(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewUserPostgresRepository(db)

	for i := 0; i < 25; i++ {
		message := "Test User " + domain.RandomString(5)
		user := &domain.User{Id: domain.UUIDv4(), Message: message}
		db.Create(user)
	}

	results, err := repo.SearchAndPaginate("Test", 2, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 10)

	cleanDB(db)
}

func setupTestDB() *gorm.DB {
	config, err := configuration.LoadConfiguration("../../..")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(config.DsnDatabase), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func cleanDB(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}
