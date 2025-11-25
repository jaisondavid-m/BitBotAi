package genai_client

import (
	"context"
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var Client *genai.Client

func InitClient() error {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	c, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return err
	}

	Client = c
	return nil
}
