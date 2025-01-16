package dependencies

import "markitos-golang-service-access/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	Delete(id *string) error
	Update(user *domain.User) error
	One(id *string) (*domain.User, error)
	List() ([]*domain.User, error)
	SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*domain.User, error)
	OneFromEmailAndPassword(email, password string) (*domain.User, error)
}
