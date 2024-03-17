package ai_test

import (
	"bytes"
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

func TestGenerateText(t *testing.T) {
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
	content, err := aiService.GenerateText("test prompt")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if content != expectedContent {
		t.Errorf("Expected %v, got %v", expectedContent, content)
	}

	assert.Equal(expectedContent, content)
}
