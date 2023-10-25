package model

type RegisterRequest struct {
	FirstName   string `json:"first_name" validate:"required,alpha" example:"Kharisma"`
	LastName    string `json:"last_name" validate:"omitempty,alpha" example:"Januar"`
	Role        string `json:"role" validate:"required,oneof=admin manager user" example:"user"`
	Email       string `json:"email" validate:"required,email" example:"kharisma.januar@gmail.com"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric,min=10" example:"08123456789"`
	Password    string `json:"password" validate:"required,min=8" example:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"kharisma.januar@gmail.com"`
	Password string `json:"password" validate:"required,min=8" example:"password"`
}
