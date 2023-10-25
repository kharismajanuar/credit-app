package model

type GetDetailUserRequest struct {
	UserID uint `params:"user_id" validate:"required,numeric" example:"1"`
}

type UpdateUserRequest struct {
	UserID      uint   `params:"user_id" validate:"required,numeric" example:"1" swaggerignore:"true"`
	Email       string `json:"email" validate:"required,email" example:"kharisma.januar@gmail.com"`
	Password    string `json:"password" validate:"required,min=8" example:"password"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric,min=10" example:"08123456789"`
	FirstName   string `json:"first_name" validate:"required,alpha" example:"Kharisma"`
	LastName    string `json:"last_name" validate:"omitempty,alpha" example:"Januar"`
	Address     string `json:"address" example:"Kp Jatilarangan"`
	Gender      string `json:"gender" validate:"omitempty,oneof=male female" example:"male"`
	DateOfBirth string `json:"date_of_birth" example:"13/01/1998"`
	Role        string `json:"role" validate:"required,oneof=admin manager user" example:"user"`
}

type DeleteUserRequest struct {
	UserID uint `params:"user_id" validate:"required,numeric" example:"1"`
}
