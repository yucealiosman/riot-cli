package cmd

import (
	"github.com/spf13/cobra"
)

var (
	token string
)

// setTokenCmd represents the token-set command
var setTokenCmd = &cobra.Command{
	Use:   "token-set",
	Short: "Set your riot api token",

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
