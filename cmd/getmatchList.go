package cmd

import (
	"encoding/json"
	"riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	matchSumName string
)

func init() {
	rootCmd.AddCommand(getmatchListCmd)

	getmatchListCmd.Flags().StringVarP(
		&matchSumName, "name", "n", "", "summoner name",
	)
	getmatchListCmd.MarkFlagRequired("name")

}

// getmatchListCmd represents the getmatchList command
var getmatchListCmd = &cobra.Command{
	Use:   "get-match-list",
	Short: "Get match list by summoner name ",

	Run: getMatchList,
}

func getMatchList(cmd *cobra.Command, args []string) {

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
	showResult(string(b))
}
