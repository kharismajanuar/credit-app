package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Username string
	Name     string
	Password string
	Port     string
	Host     string
	TimeZone string
}

type PostgresDatabase struct {
	*sqlx.DB
}

func InitPostgresql(config PostgresConfig) (database *PostgresDatabase, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", config.Host, config.Username, config.Password, config.Name, config.Port, config.TimeZone)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	database = &PostgresDatabase{
		db,
	}

	return database, nil
}
