package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"ai-summary-service/internal/handlers"
	"ai-summary-service/internal/middleware"
)

func main() {
	// Pretty logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No .env file found, using system env vars")
	}

	app := fiber.New()

	// CORS enabled
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://ai-backend-course.github.io",
		AllowMethods: "GET,POST,OPTIONS",
	}))

	// Global middleware
	app.Use(middleware.MetricsMiddleware)
	app.Use(middleware.Logger)

	// Health route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "ai-summary",
		})
	})

	app.Post("/summary", handlers.Summary)

	// Metrics route
	app.Get("/metrics", func(c *fiber.Ctx) error {
		return c.JSON(middleware.GetMetrics())
	})

	log.Info().Msg("AI Summary Service running on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
