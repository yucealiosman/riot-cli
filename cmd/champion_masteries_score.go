package cmd

import (
	"encoding/json"

	"github.com/yucealiosman/riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	name string
)

// championMasteriesScoreCmd represents the champion-masteries-score command
var championMasteriesScoreCmd = &cobra.Command{
	Use:   "champion-masteries-score",
	Short: "Champion masteries score by summoner name",
	Run:   championMasteriesScore,
}

func init() {
	rootCmd.AddCommand(championMasteriesScoreCmd)

	championMasteriesScoreCmd.Flags().StringVarP(
		&name, "name", "n", "", "summoner name",
	)
	championMasteriesScoreCmd.MarkFlagRequired("name")

}

func championMasteriesScore(cmd *cobra.Command, args []string) {
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
