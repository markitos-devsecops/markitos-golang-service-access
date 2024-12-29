package domain

type UserRepository interface {
	Create(user *User) error
	Delete(id *string) error
	Update(user *User) error
	One(id *string) (*User, error)
	List() ([]*User, error)
	SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*User, error)
}
