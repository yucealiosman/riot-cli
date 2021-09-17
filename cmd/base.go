package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func newPromtForToken() promptui.Prompt {
	return promptui.Prompt{
		Label: "Enter Riot API Token Please. You can view more details at https://developer.riotgames.com/docs/portal#_getting-started",
		Templates: &promptui.PromptTemplates{
			Valid:   "{{ . | cyan }} ",
			Success: "{{ . | blue }} ",
		},
	}
}
func setRiotToken(prompt promptui.Prompt) string {
	token, err := prompt.Run()

	if err != nil {
		showErrorAndExit(err)

	}

	fmt.Printf("You entered %q\n", token)
	setTokenInConfig(token)

	return token

}

func setTokenInConfig(token string) {
	configHandler.WriteLocalConfigFile("riotToken", token)
}

func newPromtForRegion() PromtSelect {
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
	return &prompt
}

func setRegion(prompt PromtSelect) string {
	_, region, err := prompt.Run()

	if err != nil {
		showErrorAndExit(err)
	}

	configHandler.WriteLocalConfigFile("region", region)
	fmt.Printf("You choose %q\n", region)
	return region
}

func showErrorAndExit(err error) {
	color.Red("An error has occured %v\n", err)
	os.Exit(1)
}

func showResult(cmd *cobra.Command, result string) {
	cmd.Print(result)
}
