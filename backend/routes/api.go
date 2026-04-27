package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pbleee/backend/controllers"
	"github.com/pbleee/backend/middleware"
)

func SetUp(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Regis)
	api.Post("/login", controllers.Login)
	api.Get("/course", controllers.GetCourse)

	api.Use(middleware.AuthMiddleware)

	api.Post("/course", controllers.CreateCourse)
	api.Put("/course/:id", controllers.UpdateCourse)
	api.Delete("/course/:id", controllers.DelCourse)
}
