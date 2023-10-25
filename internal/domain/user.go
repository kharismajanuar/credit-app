package domain

import (
	"time"

	"github.com/kharismajanuar/credit-app/internal/user/model"
)

type User struct {
	Users
	UserDetails
}

type Users struct {
	ID          uint
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type UserDetails struct {
	ID          uint
	UserID      uint
	FirstName   string
	LastName    string
	Address     string
	Gender      string
	DateOfBirth time.Time
	ImgKtp      []byte
	Role        string
}

type UserService interface {
	GetListUser() ([]*User, error)
	GetDetailUser(userID uint) (*User, error)
	UpdateUser(request *model.UpdateUserRequest) error
	DeleteUser(userID uint) error
}

type UserRepository interface {
	GetListUser() ([]*User, error)
	GetDetailUser(userID uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}
