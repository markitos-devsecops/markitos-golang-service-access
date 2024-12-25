package database

import (
	"fmt"
	"markitos-golang-service-access/internal/domain"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserPostgresRepository) Delete(id *string) error {
	return r.db.Delete(&domain.User{}, "id = ?", *id).Error
}

func (r *UserPostgresRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserPostgresRepository) One(id *string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", *id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepository) List() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserPostgresRepository) SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*domain.User, error) {
	offset := (pageNumber - 1) * pageSize
	var users []*domain.User
	if err := r.db.Where("message ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).
		Order("message").
		Limit(pageSize).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
