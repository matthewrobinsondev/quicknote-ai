/*
Copyright © 2024 Matthew Robinson <matthewrobinsondev@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
