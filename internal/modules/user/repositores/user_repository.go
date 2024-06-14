package repositores

import (
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}

}

func (userRepository *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User
	userRepository.DB.Create(&user).Scan(&newUser)
	return newUser
}