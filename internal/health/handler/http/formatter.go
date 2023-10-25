package http

import "github.com/kharismajanuar/credit-app/internal/domain"

type HealthCheckFormatter struct {
	AppFormatter      AppFormatter      `json:"app"`
	DatabaseFormatter DatabaseFormatter `json:"database"`
	Status            string            `json:"status"`
}

type AppFormatter struct {
	Name string `json:"name"`
}

type DatabaseFormatter struct {
	Driver string `json:"driver"`
	Status string `json:"status"`
}

func formatHealthCheck(entity *domain.Health) *HealthCheckFormatter {
	return &HealthCheckFormatter{
		AppFormatter: AppFormatter{
			Name: entity.App.Name,
		},
		DatabaseFormatter: DatabaseFormatter{
			Driver: entity.Database.Driver,
			Status: entity.Database.Status,
		},
		Status: entity.Status,
	}
}
