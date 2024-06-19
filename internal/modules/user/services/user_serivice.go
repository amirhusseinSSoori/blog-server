package services

import (
	UserModel "blog/internal/modules/user/models"
	UserResponse "blog/internal/modules/user/reponses"
	UserRepository "blog/internal/modules/user/repositores"
	auth "blog/internal/modules/user/request/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {

	return &UserService{
		userRepository: UserRepository.New(),
	}

}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user UserModel.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hasing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashPassword)

	newUser := userService.userRepository.Create(user)
	if newUser.ID == 0 {
		return response, errors.New("error on Creating User")

	}

	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	if user.ID != 0 {
		return true
	}
	return false
}
