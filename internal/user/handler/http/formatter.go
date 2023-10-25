package http

import "github.com/kharismajanuar/credit-app/internal/domain"

type ListUserFormatter struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Role        string `json:"role"`
}

func formatListUser(listUser []*domain.User) []*ListUserFormatter {
	listUserFormatted := []*ListUserFormatter{}

	for _, v := range listUser {
		user := ListUserFormatter{
			ID:          v.Users.ID,
			Email:       v.Users.Email,
			PhoneNumber: v.Users.PhoneNumber,
			FirstName:   v.UserDetails.FirstName,
			LastName:    v.UserDetails.LastName,
			Address:     v.UserDetails.Address,
			Gender:      v.UserDetails.Gender,
			DateOfBirth: v.UserDetails.DateOfBirth.Format("02/01/2006"),
			Role:        v.UserDetails.Role,
		}

		listUserFormatted = append(listUserFormatted, &user)
	}

	return listUserFormatted
}

type DetailUserFormatter struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Role        string `json:"role"`
}

func formatDetailUser(user *domain.User) *DetailUserFormatter {
	return &DetailUserFormatter{
		ID:          user.Users.ID,
		Email:       user.Users.Email,
		PhoneNumber: user.Users.PhoneNumber,
		FirstName:   user.UserDetails.FirstName,
		LastName:    user.UserDetails.LastName,
		Address:     user.UserDetails.Address,
		Gender:      user.UserDetails.Gender,
		DateOfBirth: user.UserDetails.DateOfBirth.Format("02/01/2006"),
		Role:        user.UserDetails.Role,
	}
}
