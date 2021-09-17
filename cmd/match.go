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
	rootCmd.AddCommand(matchCmd)
	matchCmd.Flags().StringVarP(
		&matchId, "match-id", "m", "", "match id",
	)
	matchCmd.MarkFlagRequired("match_id")

}

// matchCmd represents the match command
var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Match detail by match id",

	Run: match,
}

func match(cmd *cobra.Command, args []string) {

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
