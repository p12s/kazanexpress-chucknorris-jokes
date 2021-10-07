package cmd

import (
	"fmt"
	"github.com/p12s/kazanexpress-chucknorris-jokes/internal"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Getting a random joke",
	Long: `This command gets a random joke from the service 
https://api.chucknorris.io 
and displays it on the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		joke, err := internal.GetRandom()
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		fmt.Println(joke.Value)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
