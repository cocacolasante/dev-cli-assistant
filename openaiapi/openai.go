package openaiapi

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type AiRequest struct {
	Content string
}

func NewAiRequest(content string) *AiRequest{
	return &AiRequest{
		Content: content,
	}
}

func (a *AiRequest) ApiCall() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("OPEN_AI_KEY")
	client := openai.NewClient(apiKey)
	
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: a.Content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}


