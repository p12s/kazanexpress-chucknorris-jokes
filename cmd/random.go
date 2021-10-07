package cmd

import (
	"github.com/p12s/kazanexpress-chucknorris-jokes/internal"
	"github.com/spf13/cobra"
)

// randomCmd @title "random" command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Getting a random joke",
	Long: `This command gets a random joke from the service 
https://api.chucknorris.io 
and displays it on the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.RandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
