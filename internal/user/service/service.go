package service

import (
	"time"

	"github.com/kharismajanuar/credit-app/internal/domain"
	sharedConstant "github.com/kharismajanuar/credit-app/internal/shared/constant"
	sharedHelper "github.com/kharismajanuar/credit-app/internal/shared/helper"

	"github.com/kharismajanuar/credit-app/internal/user/model"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) DeleteUser(userID uint) error {
	user := domain.User{
		Users: domain.Users{
			ID:        userID,
			DeletedAt: time.Now(),
		},
	}

	return s.userRepository.DeleteUser(&user)
}

func (s *userService) GetDetailUser(userID uint) (*domain.User, error) {
	return s.userRepository.GetDetailUser(userID)
}

func (s *userService) GetListUser() ([]*domain.User, error) {
	return s.userRepository.GetListUser()
}

func (s *userService) UpdateUser(request *model.UpdateUserRequest) error {
	dateOfBirth, err := time.Parse(sharedConstant.DateLayout, request.DateOfBirth)
	if err != nil {
		return err
	}

	hashedPassword, err := sharedHelper.HashAndSalt(request.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Users: domain.Users{
			ID:          request.UserID,
			Email:       request.Email,
			Password:    hashedPassword,
			PhoneNumber: request.PhoneNumber,
			UpdatedAt:   time.Now(),
		},
		UserDetails: domain.UserDetails{
			UserID:      request.UserID,
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			Address:     request.Address,
			Gender:      request.Gender,
			DateOfBirth: dateOfBirth,
			Role:        request.Role,
		},
	}

	return s.userRepository.UpdateUser(&user)
}
