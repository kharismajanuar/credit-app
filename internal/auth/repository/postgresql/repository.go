package postgresql

import (
	"database/sql"

	"github.com/kharismajanuar/credit-app/infrastructure/database"
	"github.com/kharismajanuar/credit-app/internal/auth/model"
	"github.com/kharismajanuar/credit-app/internal/auth/model/mapper"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

type authRepository struct {
	db *database.PostgresDatabase
}

func NewAuthRepository(db *database.PostgresDatabase) domain.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) GetDetailUser(email string) (*domain.Auth, error) {
	query :=
		`
	SELECT
		a.id,
		a.email,
		a.password,
		b.role
	FROM
		public.users a,
		public.user_details b
	WHERE
		a.email = $1
		AND a.id = b.user_id
		AND a.deleted_at IS NULL
	`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, err
	}

	var authModel model.Auth
	err = stmt.Get(&authModel, email)
	if err != nil {
		return nil, err
	}

	auth := mapper.ModelToEntity(&authModel)

	return &auth, nil
}

func (r *authRepository) Register(auth *domain.Auth) error {
	authModel := mapper.EntityToModel(auth)

	tx, err := r.db.DB.Beginx()
	if err != nil {
		return err
	}

	queryInsertUsers :=
		`
		INSERT INTO public.users (
			email,
			phone_number,
			password,
			created_at,
			updated_at
		) VALUES (
			:email,
			:phone_number,
			:password,
			:created_at,
			:updated_at
		) RETURNING id
	`

	txStmt, err := tx.PrepareNamed(queryInsertUsers)
	if err != nil {
		return err
	}

	var userID int64
	err = txStmt.QueryRowx(authModel.Users).Scan(&userID)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	authModel.UserDetails.UserID = sql.NullInt64{Valid: true, Int64: userID}

	queryInsertDetailUsers :=
		`
		INSERT INTO public.user_details (
			user_id,
			first_name,
			last_name,
			role
		) VALUES (
			:user_id,
			:first_name,
			:last_name,
			:role
		)
	`

	txStmt, err = tx.PrepareNamed(queryInsertDetailUsers)
	if err != nil {
		return err
	}

	_, err = txStmt.Exec(authModel.UserDetails)
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
