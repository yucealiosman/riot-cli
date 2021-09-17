package cmd

import (
	"encoding/json"

	"github.com/yucealiosman/riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	gameSumName string
)

// currentGameCmd represents the current-game command
var currentGameCmd = &cobra.Command{
	Use:   "current-game",
	Short: "Current game by summoner name",

	Run: currentGame,
}

func init() {
	rootCmd.AddCommand(currentGameCmd)
	currentGameCmd.Flags().StringVarP(
		&gameSumName, "name", "n", "", "summoner name",
	)
	currentGameCmd.MarkFlagRequired("name")

}

func currentGame(cmd *cobra.Command, args []string) {
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
	showResult(cmd, string(b))
}
