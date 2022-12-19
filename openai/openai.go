package openai

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"os"
)

func SendToGPT(msg string) string {
	var OpenAi_Key string = os.Getenv("OpenAi_Key") // you can get it from https://beta.openai.com/account/api-keys
	ctx := context.Background()

	c := gogpt.NewClient(OpenAi_Key)

	req := gogpt.CompletionRequest{
		Model:     "ada",
		MaxTokens: 5,
		Prompt:    msg,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		log.Println("Gagal mengirimkan pesan ke GPT", err)
	}
	fmt.Println(resp.Choices[0].Text)

	return msg
}
