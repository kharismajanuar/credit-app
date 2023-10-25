package model

type CreateCreditRequest struct {
	UserID      uint    `validate:"required,numeric" swaggerignore:"true"`
	Name        string  `json:"name" validate:"required" example:"Kredit Cicilan KPR"`
	Tenor       string  `json:"tenor" validate:"required" example:"6 bulan"`
	TotalCredit float64 `json:"total_credit" validate:"required,numeric" example:"100000000"`
}

type GetListCreditRequest struct {
	UserID uint `params:"user_id" validate:"required,numeric" example:"1" swaggerignore:"true"`
}

type UpdateStatusCreditRequest struct {
	CreditID uint   `params:"credit_id" validate:"required,numeric" example:"1" swaggerignore:"true"`
	Status   string `json:"status" validate:"required,oneof=waiting processed ongoing done" example:"processed"`
}

type CreatePaymentRequest struct {
	CreditID uint    `params:"credit_id" validate:"required,numeric" example:"1" swaggerignore:"true"`
	Amount   float64 `json:"amount" validate:"required,numeric" example:"1000000"`
}

type GetListPaymentRequest struct {
	CreditID uint `params:"credit_id" validate:"required,numeric" example:"1" swaggerignore:"true"`
}

type UpdateStatusPaymentRequest struct {
	PaymentID uint   `params:"payment_id" validate:"required,numeric" example:"1" swaggerignore:"true"`
	Status    string `json:"status" validate:"required,oneof=success failed" example:"success"`
}
