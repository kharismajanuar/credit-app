package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kharismajanuar/credit-app/infrastructure/database"
	"github.com/kharismajanuar/credit-app/internal/credit/model"
	"github.com/kharismajanuar/credit-app/internal/credit/model/mapper"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

type creditRepository struct {
	db *database.PostgresDatabase
}

func NewCreditRepository(db *database.PostgresDatabase) domain.CreditRepository {
	return &creditRepository{
		db: db,
	}
}

func (r *creditRepository) CreateCredit(credit *domain.Credit) error {
	creditModel := mapper.CreditEntityToModel(credit)

	query :=
		`
		INSERT INTO public.credits (
			user_id,
			credit_type,
			name,
			tenor,
			total_credit,
			status,
			created_at,
			updated_at
		) VALUES (
			:user_id,
			'QORD',
			:name,
			:tenor,
			:total_credit,
			'waiting',
			:created_at,
			:updated_at
		)
		`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(creditModel)
	if err != nil {
		return err
	}

	return nil
}

func (r *creditRepository) CreatePayment(payment *domain.Payment) error {
	paymentModel := mapper.PaymentEntityToModel(payment)

	tx, err := r.db.DB.Beginx()
	if err != nil {
		return err
	}

	queryInsertPayments :=
		`
		INSERT INTO public.payments (
			credit_id,
			amount,
			status,
			created_at,
			updated_at
		) VALUES (
			:credit_id,
			:amount,
			'success',
			:created_at,
			:updated_at
		)
	`

	txStmt, err := tx.PrepareNamed(queryInsertPayments)
	if err != nil {
		return err
	}

	_, err = txStmt.Exec(paymentModel)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	queryGetTotalTransaction :=
		`
		SELECT
			id,
			total_transaction,
			total_credit
		FROM public.credits
		WHERE id = $1 AND deleted_at IS NULL
	`

	txGetStmt, err := tx.Preparex(queryGetTotalTransaction)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	var creditModel model.Credits
	err = txGetStmt.Get(&creditModel, paymentModel.CreditID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	creditModel.TotalTransaction.Float64 += paymentModel.Amount.Float64
	creditModel.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	if creditModel.TotalTransaction.Float64 > creditModel.TotalCredit.Float64 {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		err = fmt.Errorf("the amount you pay is above total credit: %.2f", creditModel.TotalCredit.Float64)
		return err
	}

	queryUpdateCredits :=
		`
		UPDATE public.credits 
		SET 
			total_transaction = :total_transaction,
			updated_at = :updated_at
		WHERE
			id = :id
			AND deleted_at IS NULL
	`

	txStmt, err = tx.PrepareNamed(queryUpdateCredits)
	if err != nil {
		return err
	}

	_, err = txStmt.Exec(creditModel)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	return nil
}

func (r *creditRepository) GetListCredit(userID uint) ([]*domain.Credit, error) {
	query :=
		`
		SELECT
			id,
			user_id,
			credit_type,
			name,
			total_transaction,
			tenor,
			total_credit,
			status
		FROM
			public.credits
		WHERE
			user_id = $1
			AND deleted_at IS NULL
	`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, err
	}

	var listCreditModel []*model.Credits
	err = stmt.Select(&listCreditModel, userID)
	if err != nil {
		return nil, err
	}

	return mapper.ListCreditModelToEntity(listCreditModel), err
}

func (r *creditRepository) GetListPayment(creditID uint) ([]*domain.Payment, error) {
	query :=
		`
		SELECT
			id,
			credit_id,
			amount,
			status
		FROM
			public.payments
		WHERE
			credit_id = $1
			AND deleted_at IS NULL
	`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, err
	}

	var listPaymentModel []*model.Payments
	err = stmt.Select(&listPaymentModel, creditID)
	if err != nil {
		return nil, err
	}

	return mapper.ListPaymentModelToEntity(listPaymentModel), err
}

func (r *creditRepository) UpdateStatusCredit(credit *domain.Credit) error {
	creditModel := mapper.CreditEntityToModel(credit)

	query :=
		`
		UPDATE public.credits
		SET
			status = :status,
			updated_at = :updated_at
		WHERE
			id = :id
			AND deleted_at IS NULL
		`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(creditModel)
	if err != nil {
		return err
	}

	return nil
}

func (r *creditRepository) UpdateStatusPayment(payment *domain.Payment) error {
	paymentModel := mapper.PaymentEntityToModel(payment)

	query :=
		`
		UPDATE public.payments
		SET
			status = :status,
			updated_at = :updated_at
		WHERE
			id = :id
			AND deleted_at IS NULL
		`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(paymentModel)
	if err != nil {
		return err
	}

	return nil
}
