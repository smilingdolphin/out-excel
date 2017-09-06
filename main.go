package main

import (
	"github.com/fpay/gopress"
	"github.com/out-excel/controllers"
	"github.com/out-excel/services"
)

func main() {
	// create server
	s := gopress.NewServer(gopress.ServerOptions{
		Port: 3000,
	})

	// init and register services
	s.RegisterServices(
		services.NewExcelService(),
	)

	// register middlewares
	s.RegisterGlobalMiddlewares(
		gopress.NewLoggingMiddleware("global", s.Logger),
	)

	// init and register controllers
	s.RegisterControllers(
		controllers.NewUsersController(),
		controllers.NewExcelController(),
		controllers.NewExcel2Controller(),
	)

	s.Start()
}
