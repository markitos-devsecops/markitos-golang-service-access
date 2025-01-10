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

	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
	err := repository.Create(user)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", user.Id).Error
	require.NoError(t, err)
	require.Equal(t, user.Id, result.Id)
	require.Equal(t, user.Name, result.Name)
	require.WithinDuration(t, user.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, user.UpdatedAt, result.UpdatedAt, time.Second)

	db.Delete(&result)
}

func TestSearch(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewUserPostgresRepository(db)

	randomName := domain.RandomPersonName()
	user := &domain.User{Id: domain.UUIDv4(), Name: randomName}
	db.Create(user)

	results, err := repo.SearchAndPaginate(randomName, 1, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, randomName, results[0].Name)

	cleanDB(db)
}
func TestUserDelete(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
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

	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
	db.Create(user)

	user.Name = "Updated Name"
	err := repository.Update(user)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", user.Id).Error
	require.NoError(t, err)
	require.Equal(t, "Updated Name", result.Name)
}

func TestUserOne(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
	db.Create(user)

	result, err := repository.One(&user.Id)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, user.Id, result.Id)
	require.Equal(t, user.Name, result.Name)
}

func TestUserList(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repository := database.NewUserPostgresRepository(db)

	for i := 0; i < 5; i++ {
		user := &domain.User{Id: domain.UUIDv4(), Name: domain.RandomString(10)}
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
		user := &domain.User{Id: domain.UUIDv4(), Name: domain.RandomString(10)}
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
		name := "Test User " + domain.RandomString(5)
		user := &domain.User{Id: domain.UUIDv4(), Name: name}
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
