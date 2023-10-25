package service

import (
	"github.com/kharismajanuar/credit-app/config"
	"github.com/kharismajanuar/credit-app/internal/domain"
)

type healthService struct {
	config           config.EnvConfig
	healthRepository domain.HealthRepository
}

func NewHealthService(config config.EnvConfig, healthRepository domain.HealthRepository) domain.HealthService {
	return &healthService{
		config:           config,
		healthRepository: healthRepository,
	}
}

func (s *healthService) HealthCheck() (*domain.Health, error) {
	var postgresStatus, appStatus string

	err := s.healthRepository.PostgresqlCheck()
	if err != nil {
		appStatus = "not healthy"
		postgresStatus = "disconnected"
		return nil, err
	} else {
		appStatus = "healthy"
		postgresStatus = "connected"
	}

	health := domain.Health{
		App: domain.App{
			Name: s.config.App.Name,
		},
		Database: domain.Database{
			Driver: "postgres",
			Status: postgresStatus,
		},
		Status: appStatus,
	}

	return &health, nil
}
