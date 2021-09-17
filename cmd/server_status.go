package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverStatusCmd represents the serverStatus command
var serverStatusCmd = &cobra.Command{
	Use:   "server-status",
	Short: "Show status of selected server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serverStatus called")
	},
}

func init() {
	rootCmd.AddCommand(serverStatusCmd)

}
