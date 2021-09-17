package cmd

import (
	"encoding/json"
	"riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	name string
)

// getChampionMasteriesScoreCmd represents the getChampionMasteriesScore command
var getChampionMasteriesScoreCmd = &cobra.Command{
	Use:   "get-champion-masteries-score",
	Short: "Champion masteries score by summoner name",
	Run:   getChampionMasteriesScore,
}

func init() {
	rootCmd.AddCommand(getChampionMasteriesScoreCmd)

	getChampionMasteriesScoreCmd.Flags().StringVarP(
		&name, "name", "n", "", "summoner name",
	)
	getChampionMasteriesScoreCmd.MarkFlagRequired("name")

}

func getChampionMasteriesScore(cmd *cobra.Command, args []string) {
	summoner, err := riot.GetSummonerByName(client, name)

	if err != nil {
		showErrorAndExit(err)

	}

	summonerId := summoner.Id

	championMasteryList, err := riot.GetChampionMasteriesBySummonerId(client, summonerId)

	if err != nil {
		showErrorAndExit(err)

	}

	b, err := json.MarshalIndent(*championMasteryList, "", "  ")
	if err != nil {
		showErrorAndExit(err)

	}
	showResult(string(b))
}
