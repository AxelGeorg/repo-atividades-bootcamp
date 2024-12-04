package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var percentageByDestination = &cobra.Command{
	Use:  "percentage",
	Long: "Count the number of some country percentage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			count, _ := Service.AverageDestination("")
			fmt.Printf("Number of 'all' destination occurence: %.2f\n", count*100)
		} else {
			count, _ := Service.AverageDestination(args[0])
			fmt.Printf("Number of '%s' destination occurence: %.2f\n", args[0], count*100)
		}
	},
}

func init() {
	rootCmd.AddCommand(percentageByDestination)
}
