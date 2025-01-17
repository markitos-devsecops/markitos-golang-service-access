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

func TestUserRegister(t *testing.T) {
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

func TestUserDelete(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	id := domain.UUIDv4()
	user, _ := domain.NewUser(id, domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
	db.Create(user)

	err := repository.Delete(&user.Id)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", id).Error
	require.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUserUpdateMe(t *testing.T) {
	db := setupTestDB()
	repository := database.NewUserPostgresRepository(db)

	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), domain.RandomEmail(), domain.RandomPassword(10))
	db.Create(user)

	user.Name = "Updated Name " + domain.RandomPersonName()
	err := repository.Update(user)
	require.NoError(t, err)

	var result domain.User
	err = db.First(&result, "id = ?", user.Id).Error
	require.NoError(t, err)
	require.Equal(t, user.Name, result.Name)
}

func TestUserMe(t *testing.T) {
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

// func TestOneFromEmail(t *testing.T) {
// 	db := setupTestDB()
// 	repo := database.NewUserPostgresRepository(db)

// 	user, _ := domain.NewUser(domain.UUIDv4(), domain.RandomPersonName(), "email@email.com", "anyPassw0d")
// 	db.Create(user)

// 	result, err := repo.OneFromEmail(user.Email, user.Password)
// 	require.NoError(t, err)
// 	require.NotNil(t, result)
// 	require.Equal(t, user.Id, result.Id)
// 	require.Equal(t, user.Name, result.Name)
// 	require.Equal(t, user.Email, result.Email)
// 	require.Equal(t, user.Password, result.Password)
// }

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
