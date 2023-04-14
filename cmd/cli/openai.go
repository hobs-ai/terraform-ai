package cli

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/PullRequestInc/go-gpt3"
	"github.com/ia-ops/terraform-ai/pkg/utils"
)

func (c *oaiClients) openaiGptCompletion(ctx context.Context, prompt strings.Builder, maxTokens *int, temp float32) (string, error) {
	resp, err := c.openAIClient.CompletionWithEngine(ctx, *openAIDeploymentName, openai.CompletionRequest{
		Prompt:      []string{prompt.String()},
		MaxTokens:   maxTokens,
		Echo:        false,
		N:           utils.ToPtr(1),
		Temperature: &temp,
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) != 1 {
		return "", fmt.Errorf("expected choices to be 1 but received: %d", len(resp.Choices))
	}

	return resp.Choices[0].Text, nil
}

func (c *oaiClients) openaiGptChatCompletion(ctx context.Context, prompt strings.Builder, maxTokens *int, temp float32) (string, error) {
	resp, err := c.openAIClient.ChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: *openAIDeploymentName,
		Messages: []openai.ChatCompletionRequestMessage{
			{
				Role:    userRole,
				Content: prompt.String(),
			},
		},
		MaxTokens:   *maxTokens,
		N:           1,
		Temperature: &temp,
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) != 1 {
		return "", fmt.Errorf("expected choices to be 1 but received: %d", len(resp.Choices))
	}

	return resp.Choices[0].Message.Content, nil
}