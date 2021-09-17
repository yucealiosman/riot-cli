package cmd

import (
	"encoding/json"

	"github.com/yucealiosman/riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	matchSumName string
)

func init() {
	rootCmd.AddCommand(matchListCmd)

	matchListCmd.Flags().StringVarP(
		&matchSumName, "name", "n", "", "summoner name",
	)
	matchListCmd.MarkFlagRequired("name")

}

// matchListCmd represents the match-list command
var matchListCmd = &cobra.Command{
	Use:   "match-list",
	Short: "Match list by summoner name",

	Run: matchList,
}

func matchList(cmd *cobra.Command, args []string) {

	summoner, err := riot.GetSummonerByName(client, matchSumName)

	if err != nil {
		showErrorAndExit(err)

	}

	accountId := summoner.AccountId

	championMasteryList, err := riot.GetMatchList(client, accountId)

	if err != nil {
		showErrorAndExit(err)

	}

	b, err := json.MarshalIndent(*championMasteryList, "", "  ")
	if err != nil {
		showErrorAndExit(err)

	}
	showResult(cmd, string(b))
}
