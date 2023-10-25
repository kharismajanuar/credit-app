package mapper

import (
	"database/sql"

	"github.com/kharismajanuar/credit-app/internal/credit/model"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

func CreditEntityToModel(entity *domain.Credit) model.Credits {
	return model.Credits{
		ID:               sql.NullInt64{Valid: true, Int64: int64(entity.ID)},
		UserID:           sql.NullInt64{Valid: true, Int64: int64(entity.UserID)},
		CreditType:       sql.NullString{Valid: true, String: entity.CreditType},
		Name:             sql.NullString{Valid: true, String: entity.Name},
		TotalTransaction: sql.NullFloat64{Valid: true, Float64: entity.TotalTransaction},
		Tenor:            sql.NullString{Valid: true, String: entity.Tenor},
		TotalCredit:      sql.NullFloat64{Valid: true, Float64: entity.TotalCredit},
		Status:           sql.NullString{Valid: true, String: entity.Status},
		CreatedAt:        sql.NullTime{Valid: true, Time: entity.CreatedAt},
		UpdatedAt:        sql.NullTime{Valid: true, Time: entity.UpdatedAt},
		DeletedAt:        sql.NullTime{Valid: true, Time: entity.DeletedAt},
	}
}

func CreditModelToEntity(model *model.Credits) domain.Credit {
	return domain.Credit{
		ID:               uint(model.ID.Int64),
		UserID:           uint(model.UserID.Int64),
		CreditType:       model.CreditType.String,
		Name:             model.Name.String,
		TotalTransaction: model.TotalTransaction.Float64,
		Tenor:            model.Tenor.String,
		TotalCredit:      model.TotalCredit.Float64,
		Status:           model.Status.String,
		CreatedAt:        model.CreatedAt.Time,
		UpdatedAt:        model.UpdatedAt.Time,
		DeletedAt:        model.DeletedAt.Time,
	}
}

func ListCreditModelToEntity(listModel []*model.Credits) []*domain.Credit {
	listEntity := []*domain.Credit{}
	for _, model := range listModel {
		entity := CreditModelToEntity(model)
		listEntity = append(listEntity, &entity)
	}
	return listEntity
}

func PaymentEntityToModel(entity *domain.Payment) model.Payments {
	return model.Payments{
		ID:        sql.NullInt64{Valid: true, Int64: int64(entity.ID)},
		CreditID:  sql.NullInt64{Valid: true, Int64: int64(entity.CreditID)},
		Amount:    sql.NullFloat64{Valid: true, Float64: entity.Amount},
		Status:    sql.NullString{Valid: true, String: entity.Status},
		CreatedAt: sql.NullTime{Valid: true, Time: entity.CreatedAt},
		UpdatedAt: sql.NullTime{Valid: true, Time: entity.UpdatedAt},
		DeletedAt: sql.NullTime{Valid: true, Time: entity.DeletedAt},
	}
}

func PaymentModelToEntity(model *model.Payments) domain.Payment {
	return domain.Payment{
		ID:        uint(model.ID.Int64),
		CreditID:  uint(model.CreditID.Int64),
		Amount:    model.Amount.Float64,
		Status:    model.Status.String,
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: model.UpdatedAt.Time,
		DeletedAt: model.DeletedAt.Time,
	}
}

func ListPaymentModelToEntity(listModel []*model.Payments) []*domain.Payment {
	listEntity := []*domain.Payment{}
	for _, model := range listModel {
		entity := PaymentModelToEntity(model)
		listEntity = append(listEntity, &entity)
	}
	return listEntity
}
