package service

import (
	"errors"
	"time"

	"github.com/kharismajanuar/credit-app/config"
	"github.com/kharismajanuar/credit-app/infrastructure/auth/jwt"
	"github.com/kharismajanuar/credit-app/internal/auth/model"
	"github.com/kharismajanuar/credit-app/internal/domain"
	sharedHelper "github.com/kharismajanuar/credit-app/internal/shared/helper"
)

type authService struct {
	config         config.EnvConfig
	authRepository domain.AuthRepository
	authManager    jwt.AuthManager
}

func NewAuthService(config config.EnvConfig, authRepository domain.AuthRepository, authManager jwt.AuthManager) domain.AuthService {
	return &authService{
		config:         config,
		authRepository: authRepository,
		authManager:    authManager,
	}
}

func (s *authService) Register(request *model.RegisterRequest) error {
	// Hash password
	hashedPwd, err := sharedHelper.HashAndSalt(request.Password)
	if err != nil {
		return err
	}

	request.Password = hashedPwd

	auth := domain.Auth{
		Users: domain.Users{
			Email:       request.Email,
			Password:    request.Password,
			PhoneNumber: request.PhoneNumber,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		UserDetails: domain.UserDetails{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Role:      request.Role,
		},
	}

	err = s.authRepository.Register(&auth)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(request *model.LoginRequest) (*domain.Auth, error) {
	auth, err := s.authRepository.GetDetailUser(request.Email)
	if err != nil {
		return nil, err
	}

	ok, err := sharedHelper.ComparePasswords(auth.Users.Password, request.Password)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("password is incorrect")
	}

	claim := jwt.JwtClaims{
		UserID: auth.Users.ID,
		Email:  auth.Users.Email,
		Role:   auth.UserDetails.Role,
	}

	token, err := s.authManager.CreateJwtToken(&claim)
	if err != nil {
		return nil, err
	}

	auth.Token = token

	return auth, nil
}
