package app

import (
	"log"

	"github.com/kharismajanuar/credit-app/config"
	"github.com/kharismajanuar/credit-app/infrastructure/auth/jwt"
	"github.com/kharismajanuar/credit-app/infrastructure/database"
	"github.com/kharismajanuar/credit-app/infrastructure/database/migration"
	authRepository "github.com/kharismajanuar/credit-app/internal/auth/repository/postgresql"
	authService "github.com/kharismajanuar/credit-app/internal/auth/service"
	creditRepository "github.com/kharismajanuar/credit-app/internal/credit/repository/postgresql"
	creditService "github.com/kharismajanuar/credit-app/internal/credit/service"
	"github.com/kharismajanuar/credit-app/internal/domain"
	healthRepository "github.com/kharismajanuar/credit-app/internal/health/repository/postgresql"
	healthService "github.com/kharismajanuar/credit-app/internal/health/service"
	userRepository "github.com/kharismajanuar/credit-app/internal/user/repository/postgresql"
	userService "github.com/kharismajanuar/credit-app/internal/user/service"
)

type Service struct {
	EnvironmentConfig config.EnvConfig
	HealthService     domain.HealthService
	AuthService       domain.AuthService
	UserService       domain.UserService
	CreditService     domain.CreditService
}

func SetupService() Service {
	// Init env configuration
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("unable to load .env file: %v", err)
	}

	// Init auth manager
	authManager := jwt.NewAuthManager(cfg)

	// Init postgresql database
	postgres, err := database.InitPostgresql(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to postgresql database: %v", err)
	}

	// Init database migration
	err = migration.MigratePostgresql(postgres)
	if err != nil {
		log.Fatalf("failed postgresql migration database: %v", err)
	}

	// Init repository
	healthRepository := healthRepository.NewHealthRepository(postgres)
	authRepository := authRepository.NewAuthRepository(postgres)
	userRepository := userRepository.NewUserRepository(postgres)
	creditRepository := creditRepository.NewCreditRepository(postgres)

	// Init service
	healthService := healthService.NewHealthService(cfg, healthRepository)
	authService := authService.NewAuthService(cfg, authRepository, *authManager)
	userService := userService.NewUserService(userRepository)
	creditService := creditService.NewCreditService(creditRepository)

	return Service{
		EnvironmentConfig: cfg,
		HealthService:     healthService,
		AuthService:       authService,
		UserService:       userService,
		CreditService:     creditService,
	}
}
