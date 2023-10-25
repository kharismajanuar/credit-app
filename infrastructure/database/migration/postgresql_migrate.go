package migration

import "github.com/kharismajanuar/credit-app/infrastructure/database"

func MigratePostgresql(postgres *database.PostgresDatabase) error {
	_, err := postgres.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
