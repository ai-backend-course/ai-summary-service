package ai

import "strings"

// GenerateMockSummary pretends to be an AI summarizer.
// It takes the first 20 words and add "...".
func GenerateMockSummary(text string) (string, error) {
	words := strings.Fields(text)

	if len(words) <= 20 {
		return text, nil
	}

	return strings.Join(words[:20], " ") + "...", nil
}
