// Set the API_KEY env var to your API key taken from ai.google.dev
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	prompts := make([]string, 60)
	for i := 1; i <= 60; i++ {
		prompts[i-1] = fmt.Sprintf("Who was the %d person to walk on the moon?", i)
	}
	for _, prompt := range prompts {
		fmt.Printf("\nHuman:\t%s\nstream:\t", prompt)
		answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk) + ".")
			return nil
		}))
		if err != nil {
			log.Fatalf("Failed to generate: %v, the prompt is:%s", err, prompt)
		}
		fmt.Printf("\nAI:\t%s\n", answer)
	}
}
