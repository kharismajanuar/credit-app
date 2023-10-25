package http

import "github.com/kharismajanuar/credit-app/internal/domain"

type LoginFormatter struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func formatLogin(auth *domain.Auth) *LoginFormatter {
	return &LoginFormatter{
		ID:    auth.Users.ID,
		Email: auth.Users.Email,
		Role:  auth.UserDetails.Role,
		Token: auth.Token,
	}
}
