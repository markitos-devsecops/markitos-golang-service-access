package domain

import (
	"errors"
)

type UserInMemoryRepository struct {
	users map[string]*User
}

func NewUserInMemoryRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{
		users: make(map[string]*User),
	}
}

func (s *UserInMemoryRepository) Create(user *User) error {
	s.users[user.Id] = user
	return nil
}

func (s *UserInMemoryRepository) Delete(id *string) error {
	if _, exists := s.users[*id]; !exists {
		return errors.New("user not found")
	}
	delete(s.users, *id)
	return nil
}

func (s *UserInMemoryRepository) One(id *string) (*User, error) {
	user, exists := s.users[*id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserInMemoryRepository) Update(user *User) error {
	existingUser, err := s.One(&user.Id)
	if err != nil {
		return err
	}
	existingUser.Name = user.Name
	existingUser.UpdatedAt = user.UpdatedAt
	return nil
}

func (s *UserInMemoryRepository) OneFromEmail(email string) (*User, error) {
	for _, user := range s.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
