package domain

import (
	"time"

	"github.com/kharismajanuar/credit-app/internal/credit/model"
)

type Credit struct {
	ID               uint
	UserID           uint
	CreditType       string
	Name             string
	TotalTransaction float64
	Tenor            string
	TotalCredit      float64
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type Payment struct {
	ID        uint
	CreditID  uint
	Amount    float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreditService interface {
	CreateCredit(request *model.CreateCreditRequest) error
	GetListCredit(userID uint) ([]*Credit, error)
	UpdateStatusCredit(request *model.UpdateStatusCreditRequest) error
	CreatePayment(request *model.CreatePaymentRequest) error
	GetListPayment(creditID uint) ([]*Payment, error)
	UpdateStatusPayment(request *model.UpdateStatusPaymentRequest) error
}

type CreditRepository interface {
	CreateCredit(credit *Credit) error
	GetListCredit(userID uint) ([]*Credit, error)
	UpdateStatusCredit(credit *Credit) error
	CreatePayment(payment *Payment) error
	GetListPayment(creditID uint) ([]*Payment, error)
	UpdateStatusPayment(payment *Payment) error
}
