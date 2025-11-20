package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	TotalRequests  int
	TotalErrors    int
	TotalLatencyMs float64
)

func MetricsMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	TotalRequests++
	TotalLatencyMs += float64(time.Since(start).Milliseconds())

	if err != nil {
		TotalErrors++
	}

	return err
}

func GetMetrics() fiber.Map {
	return fiber.Map{
		"total_requests":   TotalRequests,
		"total_errors":     TotalErrors,
		"total_latency_ms": TotalLatencyMs,
	}
}
