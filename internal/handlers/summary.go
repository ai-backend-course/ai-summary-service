package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"ai-summary-service/internal/ai"
)

// SummaryRequest defines the expected POST body.
type SummaryRequest struct {
	Text string `json:"text"`
}

// SummaryResponse defines what we return.
type SummaryResponse struct {
	Summary string `json:"summary"`
}

// Summary handles POST /summary
func Summary(c *fiber.Ctx) error {
	var req SummaryRequest

	// Parse JSON body into req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	if len(req.Text) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "text field is required",
		})
	}

	// Decide mock or real based on env var
	useMock := os.Getenv("USE_LLM_MOCK") == "false"

	var summary string
	var err error

	if useMock {
		summary, err = ai.GenerateMockSummary(req.Text)
	} else {
		summary, err = ai.GenerateOpenAISummary(req.Text)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return structured response
	return c.JSON(SummaryResponse{
		Summary: summary,
	})
}
