package cmd

import (
	"encoding/json"

	"github.com/yucealiosman/riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	matchId string
)

func init() {
	rootCmd.AddCommand(getmatchCmd)
	getmatchCmd.Flags().StringVarP(
		&matchId, "match-id", "m", "", "match id",
	)
	getmatchCmd.MarkFlagRequired("match_id")

}

// getmatchCmd represents the getmatch command
var getmatchCmd = &cobra.Command{
	Use:   "get-match",
	Short: "Get the match by match id",

	Run: getmatch,
}

func getmatch(cmd *cobra.Command, args []string) {

	championMasteryList, err := riot.GetMatch(client, matchId)

	if err != nil {
		showErrorAndExit(err)

	}

	b, err := json.MarshalIndent(*championMasteryList, "", "  ")
	if err != nil {
		showErrorAndExit(err)

	}
	showResult(string(b))
}
