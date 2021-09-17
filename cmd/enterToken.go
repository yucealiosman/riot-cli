package cmd

import (
	"github.com/spf13/cobra"
)

// enterTokenCmd represents the enterToken command
var enterTokenCmd = &cobra.Command{
	Use:   "enter-token",
	Short: "Create/recreate your Riot Api Token",

	Run: func(cmd *cobra.Command, args []string) {
		setRiotToken()

	},
}

func init() {
	rootCmd.AddCommand(enterTokenCmd)

}
