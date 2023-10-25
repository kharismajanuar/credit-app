package cmd

import (
	"fmt"
	"log"

	"github.com/kharismajanuar/credit-app/delivery/app"
	"github.com/kharismajanuar/credit-app/delivery/server"
)

func ListenHttp() {
	service := app.SetupService()
	handler := server.SetupHandlerHttp(service)

	app := server.Http(handler)
	err := app.Listen(fmt.Sprintf(":%s", service.EnvironmentConfig.App.Port))
	if err != nil {
		log.Panic(err)
	}
}
