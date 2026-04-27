package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/pbleee/backend/config"
	"github.com/pbleee/backend/routes"
)

func main() {
	config.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
		AllowMethods: []string{"GET, POST, PUT, DELETE"},
	}))
	routes.SetUp(app)

	app.Listen(":8080")
}
