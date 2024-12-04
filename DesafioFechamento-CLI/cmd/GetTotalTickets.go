package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var countByDestination = &cobra.Command{
	Use:  "destination",
	Long: "Count the number of some country occurences",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			count := Service.GetTotalTickets("")
			fmt.Printf("Number of 'all' destination occurence: %d\n", count)
		} else {
			count := Service.GetTotalTickets(args[0])
			fmt.Printf("Number of '%s' destination occurence: %d\n", args[0], count)
		}
	},
}

func init() {
	rootCmd.AddCommand(countByDestination)
}
