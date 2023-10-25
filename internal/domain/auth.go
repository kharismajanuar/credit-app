package domain

import (
	"github.com/kharismajanuar/credit-app/internal/auth/model"
)

type Auth struct {
	Users
	UserDetails
	Token string
}

type AuthService interface {
	Register(request *model.RegisterRequest) error
	Login(request *model.LoginRequest) (*Auth, error)
}

type AuthRepository interface {
	Register(auth *Auth) error
	GetDetailUser(email string) (*Auth, error)
}
