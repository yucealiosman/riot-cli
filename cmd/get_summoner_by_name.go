package cmd

import (
	"encoding/json"
	"sync"

	"github.com/yucealiosman/riot-cli/riot"

	"github.com/spf13/cobra"
)

var (
	nameList []string
)

func init() {
	rootCmd.AddCommand(getSummonersByNamesCmd)

	getSummonersByNamesCmd.Flags().StringSliceVarP(
		&nameList, "names", "n", nil, "list of summoner names",
	)
	getSummonersByNamesCmd.MarkFlagRequired("names")

}

// getSummonersByNamesCmd represents the getSummonerByName command
var getSummonersByNamesCmd = &cobra.Command{
	Use:   "get-summoners-by-names",
	Short: "Show summoner details with given name",
	Run:   getSummonerData,
}

func getSummonerData(cmd *cobra.Command, args []string) {

	results := make(chan *riot.Summoner, 2)
	errors := make(chan error, 2)

	var wg sync.WaitGroup

	for _, name := range nameList {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()
			result, err := riot.GetSummonerByName(client, name)
			if err != nil {
				errors <- err
				return
			}
			results <- result
		}(name)
	}

	wg.Wait()
	close(results)
	close(errors)

	showErrorList(errors)
	showResultList(results)

}

func showResultList(results chan *riot.Summoner) {
	for res := range results {
		b, err := json.MarshalIndent(*res, "", "  ")
		if err != nil {
			showErrorAndExit(err)

		}
		showResult(string(b))
	}
}

func showErrorList(errors chan error) {
	for err := range errors {
		showErrorAndExit(err)

	}
}
