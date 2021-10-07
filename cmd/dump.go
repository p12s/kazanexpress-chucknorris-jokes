package cmd

import (
	"github.com/p12s/kazanexpress-chucknorris-jokes/internal"
	"github.com/spf13/cobra"
)

// randomCmd @title "dump" command with flag
var (
	number int

	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Upload and save into files n random unique jokes",
		Long: `Upload n random unique jokes (n is set by the user, default value is 5) 
for each of the existing categories and save them to text files - 
one for each of the categories.`,
		Run: func(cmd *cobra.Command, args []string) {
			number, _ = cmd.Flags().GetInt("number")
			internal.Dump(number)
		},
	}
)

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().IntVarP(&number, "number", "n", 5, "number of jokes in each category")
}
