package ai

type AIService interface {
	GenerateText(prompt string, model string) (string, error)
}
