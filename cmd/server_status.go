package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/yucealiosman/riot-cli/riot"
)

// serverStatusCmd represents the serverStatus command
var serverStatusCmd = &cobra.Command{
	Use:   "server-status",
	Short: "Status of selected server",

	Run: func(cmd *cobra.Command, args []string) {
		serverStatus, err := riot.GetServerStatus(client)

		if err != nil {
			showErrorAndExit(err)

		}

		b, err := json.MarshalIndent(*serverStatus, "", "  ")
		if err != nil {
			showErrorAndExit(err)

		}
		showResult(cmd, string(b))

	},
}

func init() {
	rootCmd.AddCommand(serverStatusCmd)

}
