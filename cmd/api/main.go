package main

import (
	"github.com/TechmoNoway/golang-taskty-app-backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	envConfig := config.NewEnvConfig()

	app := fiber.New(fiber.Config{
		AppName:      "TasktyApp",
		ServerHeader: "Fiber",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	// server := app.Group("/api/v1")

	app.Listen(":" + envConfig.ServerPort)

}
