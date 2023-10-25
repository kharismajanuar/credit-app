package model

import (
	"database/sql"
)

type Credits struct {
	ID               sql.NullInt64   `db:"id"`
	UserID           sql.NullInt64   `db:"user_id"`
	CreditType       sql.NullString  `db:"credit_type"`
	Name             sql.NullString  `db:"name"`
	TotalTransaction sql.NullFloat64 `db:"total_transaction"`
	Tenor            sql.NullString  `db:"tenor"`
	TotalCredit      sql.NullFloat64 `db:"total_credit"`
	Status           sql.NullString  `db:"status"`
	CreatedAt        sql.NullTime    `db:"created_at"`
	UpdatedAt        sql.NullTime    `db:"updated_at"`
	DeletedAt        sql.NullTime    `db:"deleted_at"`
}

type Payments struct {
	ID        sql.NullInt64   `db:"id"`
	CreditID  sql.NullInt64   `db:"credit_id"`
	Amount    sql.NullFloat64 `db:"amount"`
	Status    sql.NullString  `db:"status"`
	CreatedAt sql.NullTime    `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"updated_at"`
	DeletedAt sql.NullTime    `db:"deleted_at"`
}
