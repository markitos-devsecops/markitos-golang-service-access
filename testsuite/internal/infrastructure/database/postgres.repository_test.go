package database_test

import (
	"log"
	"markitos-golang-service-access/internal/domain"
	"markitos-golang-service-access/internal/infrastructure/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	APP_BBDD_DSN string = "host=localhost user=admin password=admin dbname=markitos-golang-service-access sslmode=disable TimeZone=Europe/Madrid port=5432 sslmode=disable"
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
	db, err := gorm.Open(postgres.Open(APP_BBDD_DSN), &gorm.Config{})
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
