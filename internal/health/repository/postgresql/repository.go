package postgresql

import (
	"github.com/kharismajanuar/credit-app/infrastructure/database"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

type healthRepository struct {
	db *database.PostgresDatabase
}

func NewHealthRepository(db *database.PostgresDatabase) domain.HealthRepository {
	return &healthRepository{
		db: db,
	}
}

func (r *healthRepository) PostgresqlCheck() error {
	err := r.db.Ping()
	if err != nil {
		return err
	}
	return nil
}
