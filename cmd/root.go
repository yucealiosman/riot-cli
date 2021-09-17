package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/yucealiosman/riot-cli/pkg"
	"github.com/yucealiosman/riot-cli/util"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "root-cli",
	Short: "riot is a CLI tool for Riot League of Legends",
}

var client *pkg.Client

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
	util.InitConfig()
	region := viper.GetString("Region")
	if region == "" {
		region = setRegion()

	}
	token := viper.GetString("riottoken")
	if token == "" {
		setRiotToken()

	}
	err := setClient(region)
	if err != nil {
		showErrorAndExit(err)

	}

}

func setClient(region string) error {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	riotUrl := viper.GetString("riotUrl")
	var err error
	client, err = pkg.NewClient(riotUrl, httpClient)
	return err
}

func setRegion() string {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }} ",
		Active:   "{{ . | green }} ",
		Inactive: "{{ . | red }} ",
		Selected: "{{ . | bold }} ",
		Help:     "{{ . | blue }} ",
		Details:  "{{ . | yellow }} ",
	}
	prompt := promptui.Select{
		Label:     "Select a Region Please",
		Items:     []string{"americas", "asia", "europe"},
		Templates: templates,
	}

	_, region, err := prompt.Run()

	if err != nil {
		showErrorAndExit(err)

	}

	util.WriteLocalConfigFile("region", region)
	fmt.Printf("You choose %q\n", region)
	return region
}
