package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/matthewrobinsondev/quicknote-ai/pkg/ai"
	"github.com/matthewrobinsondev/quicknote-ai/pkg/markdown"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var model string
var thought string
var aiService ai.AIService

const openAIURL = "https://api.openai.com/v1/completions"

type OpenAIRequest struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief summary of the note",
	Long:  "A brief summary of the note which will turn into an obsidian note",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := aiService.GenerateText(thought, model)
		if err != nil {
			return
		}

		fmt.Println(`Creating Note:
`, content)
		markdown.SaveMarkdown(thought, content)
	},
}

func init() {
	configDir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println("Error getting config directory:", err)
		return
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(fmt.Sprintf("%s/quicknote-ai", configDir))

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("Errored reading config:", err)
		return
	}

	if !viper.IsSet("openai_api_key") {
		fmt.Println("Please add your openai_api_key to the config")
		return
	}

	client := &http.Client{}

	aiService = ai.NewOpenAIService(viper.GetString("openai_api_key"), client)

	noteCmd.Flags().StringVarP(&thought, "thought", "t", "", "No thoughts wise guy?")
	noteCmd.MarkFlagRequired("thought")

	noteCmd.Flags().StringVarP(&model, "model", "m", "gpt-3.5-turbo", "Fancy a specific model?")

	rootCmd.AddCommand(noteCmd)
}
