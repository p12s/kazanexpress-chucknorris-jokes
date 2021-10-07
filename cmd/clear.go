package cmd

import (
	"github.com/p12s/kazanexpress-chucknorris-jokes/internal"

	"github.com/spf13/cobra"
)

// randomCmd @title "clear" command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all jokes files",
	Long:  `If need to remove all created files - use this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Clear()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
