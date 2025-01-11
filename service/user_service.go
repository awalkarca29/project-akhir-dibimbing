package service

import (
	"project-akhir-awal/entity"
	"project-akhir-awal/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService interface {
	RegisterUser(input RegisterUserInput) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(input RegisterUserInput) (entity.User, error) {
	user := entity.User{}
	user.Name = input.Name
	user.Email = input.Email

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.RoleId = 2

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
