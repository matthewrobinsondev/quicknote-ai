package markdown

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func SaveMarkdown(title string, content string) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Error loading home directory", err)
	}

	if !viper.IsSet("notes_directory") {
		fmt.Println("Please add notes_directory to your config")
		return
	}

	notesDirectory := viper.GetString("notes_directory")

	filename := fmt.Sprintf("%s/%s/%s.md", homeDir, notesDirectory, title)
	err = os.WriteFile(filename, []byte(content), 0644)

	if err != nil {
		fmt.Println("Error saving markdown file:", err)
		return
	}

	fmt.Println("Markdown file saved successfully:", filename)
}
