package server

import (
	"github.com/kharismajanuar/credit-app/delivery/app"
	auth "github.com/kharismajanuar/credit-app/internal/auth/handler/http"
	credit "github.com/kharismajanuar/credit-app/internal/credit/handler/http"
	heatlh "github.com/kharismajanuar/credit-app/internal/health/handler/http"
	user "github.com/kharismajanuar/credit-app/internal/user/handler/http"
)

type HandlerHttp struct {
	healthHandler heatlh.HealthHandler
	authHandler   auth.AuthHandler
	userHandler   user.UserHandler
	creditHandler credit.CreditHandler
}

func SetupHandlerHttp(service app.Service) HandlerHttp {
	return HandlerHttp{
		healthHandler: *heatlh.NewHealthHandler(service.HealthService),
		authHandler:   *auth.NewAuthHandler(service.AuthService),
		userHandler:   *user.NewUserHandler(service.UserService),
		creditHandler: *credit.NewCreditHandler(service.CreditService),
	}
}
