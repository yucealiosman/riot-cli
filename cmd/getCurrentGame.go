package cmd

import (
	"encoding/json"
	"riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	gameSumName string
)

// getCurrentGameCmd represents the getCurrentGame command
var getCurrentGameCmd = &cobra.Command{
	Use:   "get-current-game",
	Short: "Get current game by summoner name",

	Run: getCurrentGame,
}

func init() {
	rootCmd.AddCommand(getCurrentGameCmd)
	getCurrentGameCmd.Flags().StringVarP(
		&gameSumName, "name", "n", "", "summoner name",
	)
	getCurrentGameCmd.MarkFlagRequired("name")

}

func getCurrentGame(cmd *cobra.Command, args []string) {
	summoner, err := riot.GetSummonerByName(client, gameSumName)

	if err != nil {
		showErrorAndExit(err)

	}

	summonerId := summoner.Id

	currentGame, err := riot.GetCurrentGameBySummonerId(client, summonerId)

	if err != nil {
		showErrorAndExit(err)

	}

	b, err := json.MarshalIndent(*currentGame, "", "  ")
	if err != nil {
		showErrorAndExit(err)

	}
	showResult(string(b))
}
