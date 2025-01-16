package domain

import (
	"errors"
	"strings"
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

func (s *UserInMemoryRepository) List() ([]*User, error) {
	var result []*User
	for _, value := range s.users {
		result = append(result, value)
	}

	return result, nil
}

func (s *UserInMemoryRepository) SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*User, error) {
	var filtered []*User
	for _, user := range s.users {
		if strings.Contains(user.Name, searchTerm) {
			filtered = append(filtered, user)
		}
	}

	start := (pageNumber - 1) * pageSize
	end := start + pageSize

	if start > len(filtered) {
		return []*User{}, nil
	}

	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}

func (s *UserInMemoryRepository) OneFromEmailAndPassword(email, password string) (*User, error) {
	for _, user := range s.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
