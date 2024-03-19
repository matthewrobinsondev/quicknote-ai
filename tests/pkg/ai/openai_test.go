package ai_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/matthewrobinsondev/quicknote-ai/pkg/ai"
	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGenerateText_HandleContent(t *testing.T) {
	assert := assert.New(t)

	expectedContent := "test content"

	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"choices":[{"message":{"content":"` + expectedContent + `"}}]}`
			r := io.NopCloser(bytes.NewReader([]byte(json)))

			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	aiService := ai.NewOpenAIService("test-api-key", mockClient)
	content, err := aiService.GenerateText("test prompt", "test model")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if content != expectedContent {
		t.Errorf("Expected %v, got %v", expectedContent, content)
	}

	assert.Equal(expectedContent, content)
}

func TestGenerateText_RequestContents(t *testing.T) {
	var capturedRequest *http.Request

	expectedModel := "test model"
	expectedPrompt := "test prompt"

	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			capturedRequest = req

			jsonResponse := `{"choices":[{"message":{"content":"test content"}}]}`
			r := io.NopCloser(bytes.NewReader([]byte(jsonResponse)))

			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	aiService := ai.NewOpenAIService("test-api-key", mockClient)
	_, err := aiService.GenerateText(expectedPrompt, expectedModel)

	assert.NoError(t, err)

	requestBody, err := io.ReadAll(capturedRequest.Body)

	assert.NoError(t, err)

	var payload ai.APIPayload
	err = json.Unmarshal(requestBody, &payload)

	assert.Equal(t, expectedModel, payload.Model)
	assert.Contains(t, payload.Messages[1].Content, expectedPrompt)
}

func TestGenerateText_HttpError(t *testing.T) {
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
			}, nil
		},
	}

	aiService := ai.NewOpenAIService("test-api-key", mockClient)
	_, err := aiService.GenerateText("test prompt", "test model")

	assert.Error(t, err)
}
