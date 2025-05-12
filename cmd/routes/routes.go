package routes

import (
	"context"

	"github.com/go-template-boilerplate/cmd/controllers"
	"github.com/go-template-boilerplate/cmd/middlewares"
	"github.com/go-template-boilerplate/generated"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App, ctx context.Context, queries *generated.Queries) {
	app.Post("/auth/login", controllers.LoginController(ctx, queries))
	app.Post("/auth/register", controllers.RegisterController(ctx, queries))
	app.Post("/auth/refresh_token", controllers.RefreshToken(ctx, queries))
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(middlewares.MiddleWare())
}
