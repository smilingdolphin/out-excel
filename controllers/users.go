package controllers

import (
	"net/http"

	"github.com/fpay/gopress"
)

// UsersController
type UsersController struct {
	// Uncomment this line if you want to use services in the app
	// app *gopress.App
}

// NewUsersController returns users controller instance.
func NewUsersController() *UsersController {
	return new(UsersController)
}

// RegisterRoutes registes routes to app
// It is used to implements gopress.Controller.
func (c *UsersController) RegisterRoutes(app *gopress.App) {
	// Uncomment this line if you want to use services in the app
	// c.app = app

	app.GET("/users/sample", c.SampleGetAction)
	// app.POST("/users/sample", c.SamplePostAction)
	// app.PUT("/users/sample", c.SamplePutAction)
	// app.DELETE("/users/sample", c.SampleDeleteAction)
}

// SampleGetAction Action
// Parameter gopress.Context is just alias of echo.Context
func (c *UsersController) SampleGetAction(ctx gopress.Context) error {
	// Or you can get app from request context
	// app := gopress.AppFromContext(ctx)
	data := map[string]interface{}{"service": "ExcelService", "name": ctx.QueryParam("name")}
	return ctx.Render(http.StatusOK, "users/sample", data)
}
