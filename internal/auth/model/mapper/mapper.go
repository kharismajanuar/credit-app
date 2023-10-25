package mapper

import (
	"database/sql"

	"github.com/kharismajanuar/credit-app/internal/auth/model"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

func EntityToModel(entity *domain.Auth) model.Auth {
	return model.Auth{
		Users: model.Users{
			ID:          sql.NullInt64{Valid: true, Int64: int64(entity.Users.ID)},
			Email:       sql.NullString{Valid: true, String: entity.Users.Email},
			Password:    sql.NullString{Valid: true, String: entity.Users.Password},
			PhoneNumber: sql.NullString{Valid: true, String: entity.Users.PhoneNumber},
			CreatedAt:   sql.NullTime{Valid: true, Time: entity.Users.CreatedAt},
			UpdatedAt:   sql.NullTime{Valid: true, Time: entity.Users.UpdatedAt},
			DeletedAt:   sql.NullTime{Valid: true, Time: entity.Users.DeletedAt},
		},
		UserDetails: model.UserDetails{
			ID:          sql.NullInt64{Valid: true, Int64: int64(entity.UserDetails.ID)},
			UserID:      sql.NullInt64{Valid: true, Int64: int64(entity.UserDetails.UserID)},
			FirstName:   sql.NullString{Valid: true, String: entity.UserDetails.FirstName},
			LastName:    sql.NullString{Valid: true, String: entity.UserDetails.LastName},
			Address:     sql.NullString{Valid: true, String: entity.UserDetails.Address},
			Gender:      sql.NullString{Valid: true, String: entity.UserDetails.Gender},
			DateOfBirth: sql.NullTime{Valid: true, Time: entity.UserDetails.DateOfBirth},
			Role:        sql.NullString{Valid: true, String: entity.UserDetails.Role},
		},
	}
}

func ModelToEntity(model *model.Auth) domain.Auth {
	return domain.Auth{
		Users: domain.Users{
			ID:          uint(model.Users.ID.Int64),
			Email:       model.Users.Email.String,
			Password:    model.Users.Password.String,
			PhoneNumber: model.Users.PhoneNumber.String,
			CreatedAt:   model.Users.CreatedAt.Time,
			UpdatedAt:   model.Users.UpdatedAt.Time,
			DeletedAt:   model.Users.DeletedAt.Time,
		},
		UserDetails: domain.UserDetails{
			ID:          uint(model.UserDetails.ID.Int64),
			UserID:      uint(model.UserDetails.UserID.Int64),
			FirstName:   model.UserDetails.FirstName.String,
			LastName:    model.UserDetails.LastName.String,
			Address:     model.UserDetails.Address.String,
			Gender:      model.UserDetails.Gender.String,
			DateOfBirth: model.UserDetails.DateOfBirth.Time,
			Role:        model.UserDetails.Role.String,
		},
	}
}
