package repositores

import userModel "blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
	FindByEmail(email string) userModel.User
}
