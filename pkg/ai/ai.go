package ai

type AIService interface {
	GenerateText(prompt string) (string, error)
}
