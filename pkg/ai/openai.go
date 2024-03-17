package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OpenAIService struct {
	APIKey string
	Client HTTPClient
}

type APIPayload struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
	TopP        float64   `json:"top_p"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewOpenAIService(apiKey string, client HTTPClient) AIService {
	return &OpenAIService{
		APIKey: apiKey,
		Client: client,
	}
}

func (o *OpenAIService) GenerateText(prompt string) (string, error) {
	fmt.Println("Generating text with OpenAI:", prompt)

	payload := getPayload(prompt)

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return "", err
	}

	req, err := createRequest(payloadBytes, o.APIKey)

	resp, err := o.Client.Do(req)

	body, err := io.ReadAll(resp.Body)

	var response APIResponse

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Error parsing json:", err)
		return "", err
	}

	content := response.Choices[0].Message.Content

	return content, nil
}

func createRequest(payloadBytes []byte, apiKey string) (*http.Request, error) {
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return req, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func getPayload(prompt string) APIPayload {
	return APIPayload{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "The purpose of this is you will be provided a thought and you are to complete the markdown template and return only the updated markdown. Please breakdown the thought in the note section, provide examples if possible and provide some useful resources to continue expanding the research into the thought.\n\n---\ntags: note, unassigned\nid:\n {title}\n---\n\n# Notes\n\n# Examples\n\n# Resources\n",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 1,
		MaxTokens:   700,
		TopP:        1,
	}
}
