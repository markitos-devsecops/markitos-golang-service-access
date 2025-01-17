package dependencies

import "markitos-golang-service-access/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	Delete(id *string) error
	Update(user *domain.User) error
	One(id *string) (*domain.User, error)
	OneFromEmail(email string) (*domain.User, error)
}
