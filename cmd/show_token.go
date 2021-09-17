package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showTokenCmd represents the showToken command
var showTokenCmd = &cobra.Command{
	Use:   "show-token",
	Short: "Show Riot API token",

	Run: func(cmd *cobra.Command, args []string) {
		showResult(viper.GetString("riottoken"))
	},
}

func init() {
	rootCmd.AddCommand(showTokenCmd)

}