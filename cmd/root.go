package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/yucealiosman/riot-cli/pkg"
	"github.com/yucealiosman/riot-cli/util"

	"github.com/spf13/cobra"
)

var (
	client        pkg.Client
	regionPrompt  PromtSelect        = newPromtForRegion()
	tokenPromt    promptui.Prompt    = newPromtForToken()
	configHandler util.ConfigHandler = &util.ConfigHandle{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "riotcli",
	Short: "riot is a CLI tool for Riot League of Legends",
}

type PromtSelect interface {
	Run() (int, string, error)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configHandler.InitConfig()
	region := configHandler.GetString("region")
	if region == "" {
		region = setRegion(regionPrompt)

	}
	token := configHandler.GetString("riotToken")
	if token == "" {
		setRiotToken(tokenPromt)

	}
	if (pkg.Client{}) == client {
		err := setClient(region)
		if err != nil {
			showErrorAndExit(err)

		}
	}

}

func setClient(region string) error {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	riotUrl := configHandler.GetString("riotUrl")
	riotUrl = fmt.Sprintf(riotUrl, region)

	var err error
	client, err = pkg.NewClient(riotUrl, httpClient)
	return err
}
