package cmd

import (
	"github.com/spf13/cobra"
)

var (
	token string
)

// setTokenCmdCmd represents the setToken command
var setTokenCmd = &cobra.Command{
	Use:   "set-token",
	Short: "Set your Riot Api Token",

	Run: func(cmd *cobra.Command, args []string) {
		setTokenInConfig(token)

	},
}

func init() {
	rootCmd.AddCommand(setTokenCmd)
	setTokenCmd.Flags().StringVarP(
		&token, "token", "t", "", "riot api token",
	)
	setTokenCmd.MarkFlagRequired("token")

}
