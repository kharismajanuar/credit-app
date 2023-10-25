package domain

type Health struct {
	App      App
	Database Database
	Status   string
}

type App struct {
	Name string
}

type Database struct {
	Driver string
	Status string
}

type HealthService interface {
	HealthCheck() (*Health, error)
}

type HealthRepository interface {
	PostgresqlCheck() error
}
