package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var countByPeriod = &cobra.Command{
	Use:  "period",
	Long: "Count the number of some period occurences",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			count, _ := Service.GetCountByPeriod("")
			fmt.Printf("Number of 'all' period occurence: %d\n", count)
		} else {
			count, _ := Service.GetCountByPeriod(args[0])
			fmt.Printf("Number of '%s' period occurence: %d\n", args[0], count)
		}
	},
}

func init() {
	rootCmd.AddCommand(countByPeriod)
}
