package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client handles interactions with the AI API
type Client struct {
	endpoint string
	apiKey   string
	model    string
	client   *http.Client
}

// Options configures the AI client
type Options struct {
	Endpoint string
	APIKey   string
	Model    string
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents a request to the chat completion API
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// ChatResponse represents a response from the chat completion API
type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// New creates a new AI client
func New(opts Options) *Client {
	return &Client{
		endpoint: opts.Endpoint,
		apiKey:   opts.APIKey,
		model:    opts.Model,
		client:   &http.Client{},
	}
}

// Summarize generates a summary of the provided content using the AI model
func (c *Client) Summarize(content string, systemPrompt string) (string, error) {
	url := fmt.Sprintf("%s/chat/completions", c.endpoint)

	messages := []Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: content,
		},
	}

	reqBody := ChatRequest{
		Model:    c.model,
		Messages: messages,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no summary generated")
	}

	return chatResp.Choices[0].Message.Content, nil
}
