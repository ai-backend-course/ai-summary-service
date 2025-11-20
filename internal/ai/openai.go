package ai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var (
	ErrMissingAPIKey = fmt.Errorf("OPENAI_API_KEY is missing")
	ErrEmptyResponse = fmt.Errorf("OpenAI returned no text")
)

// GenerateOpenAISummary calls OpenAI to summarize text.
func GenerateOpenAISummary(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", ErrMissingAPIKey
	}

	client := openai.NewClient(apiKey)

	// Build prompt
	prompt := "Summarize the following text in a concise paragraph:\n\n" + text

	// Create completion
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a professional text summarization assistant.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 200,
		},
	)
	if err != nil {
		return "", err
	}

	// Validate response
	if len(resp.Choices) == 0 {
		return "", ErrEmptyResponse
	}

	return resp.Choices[0].Message.Content, nil
}
