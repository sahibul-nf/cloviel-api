package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
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

func (s *service) Login(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	// cek ketersediaan email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// cek userID user tidak 0
	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	// cek kesesuaian password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("password does not match")
	}

	return user, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	// find user by ID
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// save or update data user
	user.AvatarFile = fileLocation
	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	// jika user tidak sesuai
	if user.ID == 0 {
		return user, errors.New("no user found on that ID")
	}

	return user, nil
}
