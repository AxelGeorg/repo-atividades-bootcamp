package cmd

import (
	"desafio-cli/internal/tickets"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Service tickets.Ticket

var rootCmd = &cobra.Command{
	Use: "Challange",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	Service.SetFillTicketList()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
