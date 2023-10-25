package http

import "github.com/kharismajanuar/credit-app/internal/domain"

type ListCreditFormatter struct {
	ID               uint    `json:"id"`
	UserID           uint    `json:"user_id"`
	CreditType       string  `json:"credit_type"`
	Name             string  `json:"name"`
	TotalTransaction float64 `json:"total_transaction"`
	Tenor            string  `json:"tenor"`
	TotalCredit      float64 `json:"total_credit"`
	Status           string  `json:"status"`
}

func formatListCredit(listCredit []*domain.Credit) []*ListCreditFormatter {
	listCreditFormatted := []*ListCreditFormatter{}

	for _, v := range listCredit {
		credit := ListCreditFormatter{
			ID:               v.ID,
			UserID:           v.UserID,
			CreditType:       v.CreditType,
			Name:             v.Name,
			TotalTransaction: v.TotalTransaction,
			Tenor:            v.Tenor,
			TotalCredit:      v.TotalCredit,
			Status:           v.Status,
		}
		listCreditFormatted = append(listCreditFormatted, &credit)
	}

	return listCreditFormatted
}

type ListPaymentFormatter struct {
	ID       uint    `json:"id"`
	CreditID uint    `json:"credit_id"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
}

func formatListPayment(listPayment []*domain.Payment) []*ListPaymentFormatter {
	listPaymentFormatted := []*ListPaymentFormatter{}

	for _, v := range listPayment {
		payment := ListPaymentFormatter{
			ID:       v.ID,
			CreditID: v.CreditID,
			Amount:   v.Amount,
			Status:   v.Status,
		}
		listPaymentFormatted = append(listPaymentFormatted, &payment)
	}

	return listPaymentFormatted
}
