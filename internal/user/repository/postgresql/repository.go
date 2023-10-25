package postgresql

import (
	"github.com/kharismajanuar/credit-app/infrastructure/database"
	"github.com/kharismajanuar/credit-app/internal/domain"
	"github.com/kharismajanuar/credit-app/internal/user/model"
	"github.com/kharismajanuar/credit-app/internal/user/model/mapper"
)

type userRepository struct {
	db *database.PostgresDatabase
}

func NewUserRepository(db *database.PostgresDatabase) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) DeleteUser(user *domain.User) error {
	userModel := mapper.EntityToModel(user)

	query :=
		`
		UPDATE public.users
		SET deleted_at = :deleted_at
		WHERE id = :id
		`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userModel.Users)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetDetailUser(userID uint) (*domain.User, error) {
	query :=
		`
		SELECT
			a.id,
			a.email,
			a.phone_number,
			b.first_name,
			b.last_name,
			b.address,
			b.gender,
			b.date_of_birth,
			b.role
		FROM
			public.users a,
			public.user_details b
		WHERE
			a.id = $1
			AND a.id = b.user_id
			AND a.deleted_at IS NULL
	`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, err
	}

	var userModel model.User
	err = stmt.Get(&userModel, userID)
	if err != nil {
		return nil, err
	}

	user := mapper.ModelToEntity(&userModel)

	return &user, nil
}

func (r *userRepository) GetListUser() ([]*domain.User, error) {
	query :=
		`
		SELECT
			a.id,
			a.email,
			a.phone_number,
			b.first_name,
			b.last_name,
			b.address,
			b.gender,
			b.date_of_birth,
			b.role
		FROM
			public.users a,
			public.user_details b
		WHERE
			a.id = b.user_id
			AND a.deleted_at IS NULL
	`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, err
	}

	var listUserModel []*model.User
	err = stmt.Select(&listUserModel)
	if err != nil {
		return nil, err
	}

	return mapper.ListModelToListEntity(listUserModel), err
}

func (r *userRepository) UpdateUser(user *domain.User) error {
	userModel := mapper.EntityToModel(user)

	tx, err := r.db.DB.Beginx()
	if err != nil {
		return err
	}

	queryUpdateUsers :=
		`
		UPDATE public.users 
		SET 
			email = :email,
			password = :password,
			phone_number = :phone_number,
			updated_at = :updated_at
		WHERE
			id = :id
			AND deleted_at IS NULL
	`

	txStmt, err := tx.PrepareNamed(queryUpdateUsers)
	if err != nil {
		return err
	}

	_, err = txStmt.Exec(userModel.Users)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return err
		}

		return err
	}

	queryUpdateUserDetails :=
		`
		UPDATE public.user_details 
		SET 
			first_name = :first_name,
			last_name = :last_name,
			address = :address,
			gender = :gender,
			date_of_birth = :date_of_birth,
			role = :role
		WHERE
			user_id = :user_id
	`

	txStmt, err = tx.PrepareNamed(queryUpdateUserDetails)
	if err != nil {
		return err
	}

	_, err = txStmt.Exec(userModel.UserDetails)
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
