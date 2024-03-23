package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

func generateJobDescription(inputDescription string) (string, error) {
	llm, err := openai.New(openai.WithModel("gpt-4"), openai.WithToken("")) //enter api key here
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	prompt := fmt.Sprintf(` %s  in this format: <Position Title >, About us: about the company, Briefbrief description about job, Roles and responsibilities: (6 to 8 bullet points), Knowledge and skills that will make you successful: ( 6 to 10 bullet point), (Closing line) format it using html tags.`, inputDescription)

	content := []llms.MessageContent{
		llms.TextParts(schema.ChatMessageTypeSystem, "You are a smart HR recruiter"),
		llms.TextParts(schema.ChatMessageTypeHuman, prompt),
	}

	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error  {
		return err
	}))
	if err != nil {
		log.Fatal(err)
	}
	_ = completion

	return completion.Choices[0].Content, err
}

func main() {
	inputDescription := "Manual tester with 4-6 years of experience with latest QA software knowledge responsible for verifying, testing, and identifying issues in product releases within the standards, guidelines and requirements. company name - novelio tech"
	jobDescription, err := generateJobDescription(inputDescription)
	if err != nil {
		fmt.Println("Error generating job description:", err)
		return
	}

	fmt.Println("Generated job description:", jobDescription)
}