package service

import (
	"time"

	"github.com/kharismajanuar/credit-app/internal/credit/model"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

type creditService struct {
	creditRepository domain.CreditRepository
}

func NewCreditService(creditRepository domain.CreditRepository) domain.CreditService {
	return &creditService{
		creditRepository: creditRepository,
	}
}

func (s *creditService) CreateCredit(request *model.CreateCreditRequest) error {
	credit := domain.Credit{
		UserID:      request.UserID,
		Name:        request.Name,
		Tenor:       request.Tenor,
		TotalCredit: request.TotalCredit,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.creditRepository.CreateCredit(&credit)
}

func (s *creditService) CreatePayment(request *model.CreatePaymentRequest) error {
	payment := domain.Payment{
		CreditID:  request.CreditID,
		Amount:    request.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.creditRepository.CreatePayment(&payment)
}

func (s *creditService) GetListCredit(userID uint) ([]*domain.Credit, error) {
	return s.creditRepository.GetListCredit(userID)
}

func (s *creditService) GetListPayment(creditID uint) ([]*domain.Payment, error) {
	return s.creditRepository.GetListPayment(creditID)
}

func (s *creditService) UpdateStatusCredit(request *model.UpdateStatusCreditRequest) error {
	credit := domain.Credit{
		ID:        request.CreditID,
		Status:    request.Status,
		UpdatedAt: time.Now(),
	}

	return s.creditRepository.UpdateStatusCredit(&credit)
}

func (s *creditService) UpdateStatusPayment(request *model.UpdateStatusPaymentRequest) error {
	payment := domain.Payment{
		ID:        request.PaymentID,
		Status:    request.Status,
		UpdatedAt: time.Now(),
	}

	return s.creditRepository.UpdateStatusPayment(&payment)
}
