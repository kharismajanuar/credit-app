package model

import "database/sql"

type User struct {
	Users
	UserDetails
}

type Users struct {
	ID          sql.NullInt64  `db:"id"`
	Email       sql.NullString `db:"email"`
	Password    sql.NullString `db:"password"`
	PhoneNumber sql.NullString `db:"phone_number"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
}

type UserDetails struct {
	ID          sql.NullInt64  `db:"id"`
	UserID      sql.NullInt64  `db:"user_id"`
	FirstName   sql.NullString `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	Address     sql.NullString `db:"address"`
	Gender      sql.NullString `db:"gender"`
	DateOfBirth sql.NullTime   `db:"date_of_birth"`
	ImgKtp      []byte         `db:"img_ktp"`
	Role        sql.NullString `db:"role"`
}

type ListUser struct {
	Users
	UserDetails
}

type DetailUser struct {
	Users
	UserDetails
}
