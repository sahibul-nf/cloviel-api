package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Register(input RegisterInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Register(input RegisterInput) (User, error) {

	// mapping input ke struct user
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	// hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	// save new user to db via repository
	newUser, err := s.repository.CreateNew(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
