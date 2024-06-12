package services

import (
	UserResponse "blog/internal/modules/user/reponses"
	"blog/internal/modules/user/request/auth"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}
