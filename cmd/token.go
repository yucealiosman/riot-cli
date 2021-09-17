package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TokenCmd represents the token command
var TokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Riot api token",

	Run: func(cmd *cobra.Command, args []string) {
		showResult(viper.GetString("riottoken"))
	},
}

func init() {
	rootCmd.AddCommand(TokenCmd)

}
