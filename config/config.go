package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kharismajanuar/credit-app/infrastructure/database"
)

type EnvConfig struct {
	Database database.PostgresConfig
	App      App
	Jwt      Jwt
}

type App struct {
	Port string
	Name string
}

type Jwt struct {
	Secret string
}

func LoadEnv() (envConfig EnvConfig, err error) {
	// Loading the environment variables from '.env' file.
	err = godotenv.Load()
	if err != nil {
		return EnvConfig{}, err
	}

	config := EnvConfig{
		App: App{
			Port: os.Getenv("APP_PORT"),
			Name: os.Getenv("APP_NAME"),
		},

		Database: database.PostgresConfig{
			Username: os.Getenv("DB_USERNAME"),
			Name:     os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Host:     os.Getenv("DB_HOST"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},

		Jwt: Jwt{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}

	return config, nil
}
