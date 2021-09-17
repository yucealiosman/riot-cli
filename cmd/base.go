package cmd

import (
	"fmt"
	"os"

	"github.com/yucealiosman/riot-cli/util"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func setRiotToken() string {
	prompt := promptui.Prompt{
		Label: "Enter Riot API Token Please. You can view more details at https://developer.riotgames.com/docs/portal#_getting-started",
		Templates: &promptui.PromptTemplates{
			Valid:   "{{ . | cyan }} ",
			Success: "{{ . | blue }} ",
		},
	}

	token, err := prompt.Run()

	if err != nil {
		showErrorAndExit(err)

	}

	fmt.Printf("You entered %q\n", token)
	setTokenInConfig(token)

	return token

}

func setTokenInConfig(token string) {
	util.WriteLocalConfigFile("riotToken", token)
}

func setRegion() string {
	var regionList = []string{
		"BR1",
		"EUN1",
		"EUW1",
		"JP1",
		"KR",
		"LA1",
		"LA2",
		"NA1",
		"OC1",
		"TR1",
	}

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
		Items:     regionList,
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

func showErrorAndExit(err error) {
	color.Red("An error has occured %v\n", err)
	os.Exit(1)
}

func showResult(result string) {
	color.Cyan("%v\n", result)
}
