package routes

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/BalkanID-University/cit-2025-batch-2024-summer-internship-SubikkshaRanganathan/go-jwt/pkg/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}