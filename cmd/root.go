/*
Copyright © 2024 Matthew Robinson <matthewrobinsondev@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "quicknote-ai",
	Short: "Create an ai generated note quickly on your latest idea!",
	Long: `It's for when you are trying to create a quick note on a thought that has popped into your head without leaving your state of flow.

Quick note takes the original concept of this & shoots it off into the AI mothership & spits back out a markdown file with:
- Notes
- Examples
- Resources
into your directory of choice.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
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
}
